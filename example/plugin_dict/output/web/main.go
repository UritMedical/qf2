package main

import (
	"github.com/UritMedical/qf2/define"
	"github.com/UritMedical/qf2/example/plugin_dict"
	"github.com/UritMedical/qf2/wrapper/web"
)

func main() {
	// 加载配置
	setting := web.LoadSetting(
		map[string]define.QPlugin{
			"Dict": &plugin_dict.Plugin{},
		},
		web.RouteRule{
			PublishRule: web.PublishRule{}, // 无规则，全部暴露，后续可以通过配置文件修改
			TransRoutes: nil,               // 无转换，全部使用默认的路由，后续可以通过配置文件修改
		},
		[]web.Middlewares{apiMiddleware},
		[]web.Middlewares{msgMiddleware},
	)
	// 启动
	web.Start(setting, nil)
}

func apiMiddleware(pluginKey string, funcName string, ctx web.IContext) {

}

func msgMiddleware(pluginKey string, funcName string, ctx web.IContext) {

}
