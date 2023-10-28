package web

import (
	"github.com/UritMedical/qf2/define"
	"reflect"
	"strings"
)

func newAdapter() *adapter {
	return &adapter{
		apiHandlers: map[string]define.QApiHandler{},
	}
}

type adapter struct {
	apiHandlers map[string]define.QApiHandler
}

// RegApi
//
//	@Description: 支持API方法
//	@param name 方法名称
//	@param handler 方法指针
func (a *adapter) RegApi(name string, handler define.QApiHandler) {
	name = strings.ToLower(name)
	// 去掉root
	if strings.HasPrefix(name, "root_") {
		name = strings.Replace(name, "root_", "", 1)
	}
	a.apiHandlers[name] = handler
}

func (a *adapter) getRoutes() []route {
	routes := make([]route, 0)
	for k, _ := range a.apiHandlers {
		sp := strings.Split(k, "_")
		// 生成地址
		url := ""
		for i := 0; i < len(sp)-1; i++ {
			url += "/" + strings.ToLower(sp[i])
		}
		url = strings.Trim(url, "/")
		// 添加到字典
		routes = append(routes, route{url, strings.ToLower(sp[len(sp)-1])})
	}
	return routes
}

func (a *adapter) doApi(ctx *context) (interface{}, define.QFail) {
	url := strings.Trim(ctx.gin.FullPath(), "/")
	url = strings.Replace(url, "/", "_", -1)
	name := url + "_" + strings.ToLower(ctx.gin.Request.Method)
	if handler, ok := a.apiHandlers[name]; ok {
		return handler(ctx)
	}
	return nil, nil
}

func (a *adapter) RegSubscribe(name string, handler define.QNoticeHandler) {

}

func (a *adapter) SendNotice(topic string, payload interface{}) {

}

func (a *adapter) Invoke(name string, params []interface{}, of reflect.Type) interface{} {
	return nil
}
