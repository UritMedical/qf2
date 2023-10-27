package web

import (
	"errors"
	"fmt"
	"github.com/UritMedical/qf2/define"
	"github.com/UritMedical/qf2/utils/launcher"
	"github.com/UritMedical/qf2/utils/qerror"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Run
//
//	@Description: 启动
//	@param params
//	@param stop
func Run(setting WebParams, stop func()) {
	ginWeb := &ginWeb{
		params:   setting,
		stopFunc: stop,
	}
	launcher.Run(ginWeb.Start, ginWeb.Stop)
}

type ginWeb struct {
	params   WebParams
	stopFunc func()

	engine  *gin.Engine
	setting *setting
	adapter *adapter
}

func (gw *ginWeb) Start() {
	// 收集异常
	defer qerror.Recover(func(err string) {
		launcher.Exit()
	})

	// 加载配置
	gw.setting = loadSetting()

	// 初始化插件
	gw.initPlugin()

	// 初始化服务
	gw.engine = gin.Default()
	gw.engine.Use(gw.getCors())
	gw.initRoute()

	// 启动服务
	go func() {
		err := gw.engine.Run(fmt.Sprintf(":%d", gw.setting.Port))
		if err != nil {
			panic(err)
		}
	}()
}

func (gw *ginWeb) Stop() {
	gw.params.Bll.Stop()
}

func (gw *ginWeb) initPlugin() {
	// 创建访问器
	gw.adapter = newAdapter()

	// 初始化业务
	gw.params.Bll.Init()
	for _, initFunc := range gw.params.AppInitFunc {
		initFunc(gw.adapter)
	}
	// 绑定业务方法
	gw.params.Bll.Bind()
}

func (gw *ginWeb) initRoute() {
	for url, tp := range gw.adapter.getRoutes() {
		fullUrl := "/" + url
		if gw.setting.DefGroup != "" {
			fullUrl = "/" + gw.setting.DefGroup + "/" + url
		}
		switch tp {
		case "get":
			gw.engine.GET(fullUrl, gw.apiRequest)
		case "post":
			gw.engine.POST(fullUrl, gw.apiRequest)
		case "put":
			gw.engine.PUT(fullUrl, gw.apiRequest)
		case "delete":
			gw.engine.DELETE(fullUrl, gw.apiRequest)
		}
	}
}

func (gw *ginWeb) apiRequest(ginCtx *gin.Context) {
	// 创建上下文
	ctx := newContext(ginCtx)
	// 执行方法
	result, err := gw.adapter.doApi(ctx)
	// 返回
	if err != nil {
		if e, ok := err.(define.Error); ok {
			gw.returnError(ginCtx, e)
		} else if r, ok := err.(define.Refuse); ok {
			gw.returnRefuse(ginCtx, r)
		} else {
			gw.returnError(ginCtx, define.NewError(500, errors.New("未知的错误类型")))
		}
	} else {
		gw.returnOk(ginCtx, result)
	}
}

func (gw *ginWeb) returnOk(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "success",
		"data":   data,
	})
}

func (gw *ginWeb) returnError(ctx *gin.Context, err define.Error) {
	// 记录日志
	qerror.Write(fmt.Sprintf("\n\t%s %s %d %s %s", ctx.Request.Method, ctx.Request.URL, err.Code(), err.Desc(), err.Error()))
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"status":    http.StatusInternalServerError,
		"msg":       err.Desc(),
		"exception": err.Error(),
		"data":      err.Code(),
	})
}

func (gw *ginWeb) returnRefuse(ctx *gin.Context, err define.Refuse) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"status": http.StatusBadRequest,
		"msg":    err.Desc(),
		"data":   err.Code(),
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
