package qweb

import (
	"github.com/UritMedical/qf2/qdefine"
	"sort"
	"strings"
)

func newAdapter() *adapter {
	return &adapter{
		apiHandlers: map[string]qdefine.QApiHandler{},
	}
}

type adapter struct {
	apiHandlers    map[string]qdefine.QApiHandler
	tmpLastApiName string
}

// RegApi
//
//	@Description: 支持API方法
//	@param name 方法名称
//	@param handler 方法指针
func (a *adapter) RegApi(name string, handler qdefine.QApiHandler) {
	name = strings.ToLower(name)
	// 去掉root
	if strings.HasPrefix(name, "root_") {
		name = strings.Replace(name, "root_", "", 1)
	}
	a.apiHandlers[name] = handler
	a.tmpLastApiName = name
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
	// 排序
	sort.Slice(routes, func(i, j int) bool {
		return routes[i].Url > routes[j].Url
	})
	return routes
}

func (a *adapter) formatUrlToName(ctx *context, defGroup string) string {
	url := ctx.url
	url = strings.Replace(url, defGroup, "", 1)
	url = strings.Trim(url, "/")
	url = strings.Replace(url, "/", "_", -1)
	name := url + "_" + strings.ToLower(ctx.method)
	return name
}

func (a *adapter) doApi(ctx *context, defGroup string) (interface{}, qdefine.QFail) {
	name := a.formatUrlToName(ctx, defGroup)
	if handler, ok := a.apiHandlers[name]; ok {
		return handler(ctx)
	}
	return nil, nil
}

func (a *adapter) RegSubscribe(name string, handler qdefine.QNoticeHandler) {

}

func (a *adapter) SendNotice(topic string, payload interface{}) {

}

func (a *adapter) Invoke(funcName string, params map[string]interface{}) (interface{}, qdefine.QFail) {
	ctx := newContextByRef(funcName, params)
	if handler, ok := a.apiHandlers[funcName]; ok {
		return handler(ctx)
	}
	return nil, nil
}
