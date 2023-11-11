package qweb

import (
	"fmt"
	"github.com/UritMedical/qf2/qdefine"
	"github.com/UritMedical/qf2/utils/launcher"
	"github.com/UritMedical/qf2/utils/qdb"
	"github.com/UritMedical/qf2/utils/qerror"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

const SettingPath string = "./config/setting.yaml"

// Run
//
//	@Description: 启动
//	@param Widget
func Run(f func(), widget qdefine.QWidget) {
	ginWeb := &ginWeb{
		Widget:   widget,
		initFunc: f,
	}
	launcher.Run(ginWeb.Start, ginWeb.Stop)
}

type ginWeb struct {
	///启动参数
	Widget   qdefine.QWidget
	initFunc func()
	//gin引擎
	engine *gin.Engine
	//设置
	setting *setting
	//qf访问器
	adapter *adapter
	////中间件
	//middleware map[string]qdefine.QBll
}

func (gw *ginWeb) Start() {
	// 收集异常
	defer qerror.Recover(func(err string) {
		launcher.Exit()
	})

	// 加载配置
	gw.setting = newSetting(SettingPath)
	time.Sleep(time.Second * time.Duration(gw.setting.StartDelay))

	gw.initFunc()
	// 初始化微件
	gw.initWidget()

	// 初始化服务
	gw.engine = gin.Default()
	gw.engine.Use(gw.getCors()) // 支持跨域
	//gw.engine.Use(gw.apiMiddleware()) // 加载中间件
	gw.initRoute()

	// 保存配置
	gw.setting.Save()

	// 启动服务
	go func() {
		err := gw.engine.Run(fmt.Sprintf(":%d", gw.setting.Port))
		if err != nil {
			panic(err)
		}
	}()
}

func (gw *ginWeb) Stop() {
	for _, m := range gw.Widget.Modules {
		m.OnStop()
	}
}

func (gw *ginWeb) initWidget() {
	// 创建访问器
	gw.adapter = newAdapter()

	// 初始化Dao
	qdb.ConfigPath = SettingPath

	gw.setting.GormConfig = qdb.Settings

	//// 初始化业务
	//gw.middleware = make(map[string]qdefine.QBll)

	//if gw.Widget.ReferencesInit != nil {
	//	gw.Widget.ReferencesInit(gw.adapter)
	//}

	// 绑定业务方法
	for k, m := range gw.Widget.Modules {
		m.Reg(k+"/", gw.adapter)
	}
}

func (gw *ginWeb) initRoute() {
	for k := range gw.adapter.getApis {
		gw.engine.GET(gw.setting.DefGroup+"/"+k, gw.apiRequest)
	}
	for k := range gw.adapter.postApis {
		gw.engine.POST(gw.setting.DefGroup+"/"+k, gw.apiRequest)
	}
	for k := range gw.adapter.putApis {
		gw.engine.PUT(gw.setting.DefGroup+"/"+k, gw.apiRequest)
	}
	for k := range gw.adapter.delApis {
		gw.engine.DELETE(gw.setting.DefGroup+"/"+k, gw.apiRequest)
	}
}

func (gw *ginWeb) apiRequest(ginCtx *gin.Context) {
	// 创建上下文
	ctx := newContextByGin(ginCtx)

	ctx.route = strings.Replace(ctx.route, gw.setting.DefGroup, "", 1)
	ctx.route = strings.Trim(ctx.route, "/")

	// 前置
	head := strings.Split(ctx.route, "/")[0]
	if module, ok := gw.Widget.Modules[head]; ok {
		fail := module.OnStartInvoke(ctx.route, ctx)
		if fail != nil {
			gw.returnErr(ginCtx, fail)
			return
		}
	}

	// 执行方法
	result, fail := gw.adapter.doApi(ctx)

	// 返回
	if fail != nil {
		gw.returnErr(ginCtx, fail)
	} else {
		// 后置
		ctx.SetNewReturnValue(result)
		if module, ok := gw.Widget.Modules[head]; ok {
			module.OnEndInvoke(ctx.route, ctx)
		}
		// 返回
		gw.returnOk(ginCtx, ctx.GetReturnValue())
	}
}

func (gw *ginWeb) returnErr(ginCtx *gin.Context, fail *qdefine.QFail) {
	if fail.Err != nil {
		gw.returnError(ginCtx, fail.Err)
	} else {
		gw.returnRefuse(ginCtx, fail.Code, fail.Desc)
	}
}

func (gw *ginWeb) returnOk(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "success",
		"data":   data,
	})
}

func (gw *ginWeb) returnError(ctx *gin.Context, err error) {
	// 记录日志
	qerror.Write(fmt.Sprintf("\n\t%s %s %s", ctx.Request.Method, ctx.Request.URL, err.Error()))
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"status": http.StatusInternalServerError,
		"msg":    strings.Trim(err.Error(), " "),
	})
}

func (gw *ginWeb) returnRefuse(ctx *gin.Context, code int, desc string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"status": http.StatusBadRequest,
		"msg":    strings.Trim(desc, " "),
		"code":   code,
	})
}

func (gw *ginWeb) getCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
