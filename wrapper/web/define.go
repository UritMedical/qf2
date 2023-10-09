package web

import "github.com/UritMedical/qf2/define"

type Wrapper struct {
	engine  IEngine
	setting Setting
}

type Middlewares func(pluginKey string, funcName string, ctx IContext)

type WrapperSetting struct {
	Engine         IEngine
	Setting        Setting
	RouteRule      RouteRule
	Plugins        []define.QPlugin
	ApiMiddlewares []Middlewares
	MsgMiddlewares []Middlewares
}

type IEngine interface {
}

type IContext interface {
}

type Setting struct {
	Port          int                 // 自身服务端口
	RefPluginPort map[string]int      // 引用其他插件的服务端口
	BeSubscribed  map[string][]string // 主题被订阅的插件列表
}

type RouteRule struct {
	PublishRule PublishRule  // 路由暴露规则
	TransRoutes []TransRoute // 路由转换规则
}

type PublishRule struct {
	IsBlack bool     // 是否是黑名单
	Rules   []string // 规则列表，可以是指定路由或者正则表达式
}

type TransRoute struct {
	Source string // 默认生成的路由 例如：post:/api/xxx
	New    string // 转换后的路由 例如：get:/api/xxx
}

//func LoadSetting(args []string) Setting {
//	return Setting{}
//}

//func LoadDefaultSetting() Setting {
//	return Setting{}
//}
