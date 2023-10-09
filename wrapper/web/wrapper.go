package web

import (
	"fmt"
	"github.com/UritMedical/qf2/define"
	"github.com/UritMedical/qf2/utils/launcher"
)

func Start(setting WrapperSetting, stop func()) {
	launcher.RunEx(start, setting, stop)
}

func LoadSetting(plugins map[string]define.QPlugin, rule RouteRule, apisMiddlewares []Middlewares, msgMiddlewares []Middlewares) WrapperSetting {
	return WrapperSetting{}
}

func start(param interface{}) {
	setting := param.(WrapperSetting)
	fmt.Println(setting)
}

//
//func NewWrapper() *Wrapper {
//	return &Wrapper{}
//}
//
//func (w *Wrapper) RegEngine(engine IEngine, setting Setting) {
//	w.engine = engine
//	w.setting = setting
//}
//
//func (w *Wrapper) LoadRule(def func() RouteRule) {
//
//}
//
//func (w *Wrapper) RegPlugins(plugins []define.QPlugin) {
//
//}
//
//func (w *Wrapper) ApiMiddleware(handler func(key string, ctx IContext)) {
//
//}
//
//func (w *Wrapper) MsgMiddleware(handler func(key string, ctx IContext)) {
//
//}
//
//func (w *Wrapper) Run() {
//
//}
