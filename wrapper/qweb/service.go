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
	startParam StartParam
	engine     *gin.Engine
	setting    *setting
	adapter    *adapter
	middleware map[string]qdefine.QBllSvc
}

func (gw *ginWeb) Start() {
	// 收集异常
	defer qerror.Recover(func(err string) {
		launcher.Exit()
	})

	// 加载配置
	gw.setting = newSetting(gw.startParam.ConfigPath)

	// 初始化插件
	gw.initPlugin()

	// 初始化服务
	gw.engine = gin.Default()
	gw.engine.Use(gw.getCors())       // 支持跨域
	gw.engine.Use(gw.apiMiddleware()) // 加载中间件
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
	for _, svc := range gw.startParam.Svcs {
		svc.DaoSvc.Stop()
		svc.BllSvc.Stop()
	}
}

func (gw *ginWeb) initPlugin() {
	// 创建访问器
	gw.adapter = newAdapter()

	// 初始化Dao
	qdb.ConfigPath = gw.startParam.ConfigPath
	for _, svc := range gw.startParam.Svcs {
		svc.DaoSvc.Init()
	}
	gw.setting.GormConfig = qdb.Settings

	// 初始化业务
	gw.middleware = make(map[string]qdefine.QBllSvc)
	for _, svc := range gw.startParam.Svcs {
		svc.BllSvc.Init()
		svc.QfSvc(gw.adapter)

		// 提取该模块注册的服务前缀，生成每个模块的中间件字段
		// 用于根据url调用对应模块的中间件
		pkg := strings.Split(gw.adapter.tmpLastApiName, "_")[0]
		gw.middleware[pkg] = svc.BllSvc
	}
	if gw.startParam.ReferencesInit != nil {
		gw.startParam.ReferencesInit(gw.adapter)
	}

	// 绑定业务方法
	for _, svc := range gw.startParam.Svcs {
		svc.BllSvc.Bind()
	}
}

func (gw *ginWeb) initRoute() {
	for _, route := range gw.adapter.getRoutes() {
		fullUrl := "/" + route.Url
		if gw.setting.DefGroup != "" {
			fullUrl = "/" + gw.setting.DefGroup + "/" + route.Url
		}
		switch strings.ToLower(route.Type) {
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
	ctx := newContextByGin(ginCtx)
	// 执行方法
	result, err := gw.adapter.doApi(ctx, gw.setting.DefGroup)
	// 返回
	if err != nil {
		if e, ok := err.(qdefine.Error); ok {
			gw.returnError(ginCtx, e)
		} else if r, ok := err.(qdefine.Refuse); ok {
			gw.returnRefuse(ginCtx, r)
		} else {
			gw.returnError(ginCtx, qdefine.NewError(500, errors.New("未知的错误类型")))
		}
	} else {
		gw.returnOk(ginCtx, result)
	}
}

func (gw *ginWeb) apiMiddleware() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		// 创建上下文
		ctx := newContextByGin(ginCtx)
		name := gw.adapter.formatUrlToName(ctx, gw.setting.DefGroup)
		per := strings.Split(name, "_")[0]
		if bll, ok := gw.middleware[per]; ok {
			err := bll.Middleware(name, ctx)
			// 拒绝
			if err != nil {
				if e, ok := err.(qdefine.Error); ok {
					gw.returnError(ginCtx, e)
				} else if r, ok := err.(qdefine.Refuse); ok {
					gw.returnRefuse(ginCtx, r)
				} else {
					gw.returnError(ginCtx, qdefine.NewError(500, errors.New("未知的错误类型")))
				}
				ginCtx.Abort()
				return
			}
		}
		// 继续
		ginCtx.Next()
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
	qerror.Write(fmt.Sprintf("\n\t%s %s %d %s %s", ctx.Request.Method, ctx.Request.URL, err.Code(), err.Desc(), err.Error()))
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"status":    http.StatusInternalServerError,
		"msg":       err.Desc(),
		"exception": err.Error(),
		"data":      err.Code(),
	})
}

func (gw *ginWeb) returnRefuse(ctx *gin.Context, err qdefine.Refuse) {
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
