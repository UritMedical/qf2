package qweb

import (
	"errors"
	"fmt"
	"github.com/UritMedical/qf2/qdefine"
	"github.com/UritMedical/qf2/utils/launcher"
	"github.com/UritMedical/qf2/utils/qdb"
	"github.com/UritMedical/qf2/utils/qerror"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const SettingPath string = "./config/setting.yaml"

// Run
//
//	@Description: 启动
//	@param startParam
func Run(startParam StartParam) {
	ginWeb := &ginWeb{
		startParam: startParam,
	}
	launcher.Run(ginWeb.Start, ginWeb.Stop)
}

type ginWeb struct {
	///启动参数
	startParam StartParam
	//gin引擎
	engine *gin.Engine
	//设置
	setting *setting
	//qf访问器
	adapter *adapter
	//中间件
	middleware map[string]qdefine.QBll
}

func (gw *ginWeb) Start() {
	// 收集异常
	defer qerror.Recover(func(err string) {
		launcher.Exit()
	})

	// 加载配置
	gw.setting = newSetting(SettingPath)

	// 初始化插件
	gw.initPlugin()

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
	for _, m := range gw.startParam.Modules {
		m.QDao.Stop()
		m.QBll.Stop()
	}
}

func (gw *ginWeb) initPlugin() {
	// 创建访问器
	gw.adapter = newAdapter()

	// 初始化Dao
	qdb.ConfigPath = SettingPath
	for _, m := range gw.startParam.Modules {
		m.QDao.Init()
	}
	gw.setting.GormConfig = qdb.Settings

	// 初始化业务
	gw.middleware = make(map[string]qdefine.QBll)
	for k, m := range gw.startParam.Modules {
		m.QBll.Init()
		//m.QAdapter(gw.adapter)

		// 提取该模块注册的服务前缀，生成每个模块的中间件字段
		// 用于根据url调用对应模块的中间件
		//pkg := strings.Split(gw.adapter.tmpLastApiName, "_")[0]

		gw.middleware[k] = m.QBll
	}
	//if gw.startParam.ReferencesInit != nil {
	//	gw.startParam.ReferencesInit(gw.adapter)
	//}

	// 绑定业务方法
	for k, m := range gw.startParam.Modules {
		m.QBll.Bind(k+"/", gw.adapter)
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
	// 前置
	ctx.route = strings.Replace(ctx.route, gw.setting.DefGroup, "", 1)

	head := strings.Split(ctx.route, "/")[0]
	if bll, ok := gw.middleware[head]; ok {
		err := bll.StartInvoke(ctx.route, ctx)
		if err != nil {
			gw.returnErr(ginCtx, err)
			return
		}
	}

	// 执行方法
	result, err := gw.adapter.doApi(ctx)

	// 返回
	if err != nil {
		gw.returnErr(ginCtx, err)
	} else {
		// 后置
		ctx.SetNewReturnValue(result)
		if bll, ok := gw.middleware[head]; ok {
			bll.EndInvoke(ctx.route, ctx)
		}
		// 返回
		gw.returnOk(ginCtx, ctx.GetReturnValue())
	}
}

func (gw *ginWeb) returnErr(ginCtx *gin.Context, err qdefine.QFail) {
	if e, ok := err.(qdefine.Error); ok {
		gw.returnError(ginCtx, e)
	} else if r, ok := err.(qdefine.Refuse); ok {
		gw.returnRefuse(ginCtx, r)
	} else {
		gw.returnError(ginCtx, qdefine.NewError(qdefine.ErrorCodeOSError, errors.New("未知的错误类型")))
	}
}

func (gw *ginWeb) returnOk(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "success",
		"data":   data,
	})
}

func (gw *ginWeb) returnError(ctx *gin.Context, err qdefine.Error) {
	// 记录日志
	qerror.Write(fmt.Sprintf("\n\t%s %s %s %s %s", ctx.Request.Method, ctx.Request.URL, err.Code(), err.Desc(), err.Error()))
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"status":  http.StatusInternalServerError,
		"msg":     strings.Trim(fmt.Sprintf("%s %s", err.Desc(), err.Error()), " "),
		"errCode": err.Code(),
	})
}

func (gw *ginWeb) returnRefuse(ctx *gin.Context, err qdefine.Refuse) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"status":  http.StatusBadRequest,
		"msg":     strings.Trim(fmt.Sprintf("%s %s", err.Desc(), err.Error()), " "),
		"errCode": err.Code(),
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
