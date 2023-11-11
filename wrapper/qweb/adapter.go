package qweb

import (
	"errors"
	"fmt"
	"github.com/UritMedical/qf2/qdefine"
	"reflect"
)

func newAdapter() *adapter {
	return &adapter{
		getApis:  map[string]qdefine.QApiHandler{},
		postApis: map[string]qdefine.QApiHandler{},
		putApis:  map[string]qdefine.QApiHandler{},
		delApis:  map[string]qdefine.QApiHandler{},
	}
}

type adapter struct {
	getApis  map[string]qdefine.QApiHandler
	postApis map[string]qdefine.QApiHandler
	putApis  map[string]qdefine.QApiHandler
	delApis  map[string]qdefine.QApiHandler
}

func (a *adapter) RegGet(router string, handler qdefine.QApiHandler) {
	a.getApis[router] = handler
}

func (a *adapter) RegPost(router string, handler qdefine.QApiHandler) {
	a.postApis[router] = handler
}

func (a *adapter) RegPut(router string, handler qdefine.QApiHandler) {
	a.putApis[router] = handler
}

func (a *adapter) RegDel(router string, handler qdefine.QApiHandler) {
	a.delApis[router] = handler
}

func (a *adapter) Invoke(route, method string, params map[string]interface{}) (interface{}, *qdefine.QFail) {
	ctx := newContextByRef(route, method, params)
	switch method {
	case "Get":
		if handler, ok := a.getApis[route]; ok {
			return handler(ctx)
		}
	case "Put":
		if handler, ok := a.putApis[route]; ok {
			return handler(ctx)
		}
	case "Post":
		if handler, ok := a.postApis[route]; ok {
			return handler(ctx)
		}
	case "Del":
		if handler, ok := a.delApis[route]; ok {
			return handler(ctx)
		}
	}
	return nil, nil
}

func (a *adapter) doApi(ctx *context) (interface{}, *qdefine.QFail) {
	route := ctx.route
	method := ctx.method
	var handler qdefine.QApiHandler = nil
	switch method {
	case "POST":
		if h, ok := a.postApis[route]; ok {
			handler = h
		}
	case "DELETE":
		if h, ok := a.delApis[route]; ok {
			handler = h
		}
	case "PUT":
		if h, ok := a.putApis[route]; ok {
			handler = h
		}
	case "GET":
		if h, ok := a.getApis[route]; ok {
			handler = h
		}
	}
	if handler != nil {
		r := reflect.ValueOf(handler)
		b := r.Interface()
		fmt.Println(b)
		return handler(ctx)
	} else {
		return nil, &qdefine.QFail{Err: errors.New(fmt.Sprintf("unknown api %s", method)), Desc: "unknown api"}
	}
}

func (a *adapter) SendNotice(topic string, payload interface{}) {
	a.SendNotice(topic, payload)
}
