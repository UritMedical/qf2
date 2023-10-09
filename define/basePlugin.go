package define

import (
	"errors"
	"fmt"
)

type BasePlugin struct {
	bllList      []QBll
	apis         map[string]ApiHandler
	subscribes   map[string]NoticeHandler
	pluginInvoke func(plugin string, route string, params ...interface{}) (interface{}, error)
	pluginNotice func(topic string, params ...interface{}) error
	//mqClient     *mqtt.Client
}

func (b *BasePlugin) RegBll(bll ...QBll) {
	if b.apis == nil {
		b.apis = map[string]ApiHandler{}
	}
	if b.subscribes == nil {
		b.subscribes = map[string]NoticeHandler{}
	}
	for _, v := range bll {
		for k, v := range v.Apis() {
			b.apis[k] = v
		}
		for k, v := range v.Subscribes() {
			b.subscribes[k] = v
		}
	}
	b.bllList = bll
}

func (b *BasePlugin) BuildDocument() {
	for _, bll := range b.bllList {
		config := bll.GetConfig()
		fmt.Println(config)
	}
}

func (b *BasePlugin) BindWrapperInvoke(invoke func(plugin string, route string, params ...interface{}) (interface{}, error)) {
	b.pluginInvoke = invoke
}

func (b *BasePlugin) BindWrapperNotice(notice func(topic string, params ...interface{}) error) {
	b.pluginNotice = notice
}

func (b *BasePlugin) Invoke(route string, params ...interface{}) (interface{}, error) {
	// 调用插件内部的业务方法
	if handler, ok := b.apis[route]; ok {
		if handler == nil {
			return nil, errors.New("")
		}
		return handler(params...)
	}
	if handler, ok := b.subscribes[route]; ok {
		if handler == nil {
			return nil, errors.New("")
		}
		handler(params...)
		return nil, nil
	}
	// 调用其他插件的业务方法
	if b.pluginInvoke == nil {
		return nil, errors.New("plugin invoke not init")
	}
	return b.pluginInvoke("", route, params)
}

func (b *BasePlugin) SendNotice(topic string, params ...interface{}) error {
	//// 发送MQ消息。给前端或者其他有需要的客户端使用
	//if b.mqClient != nil {
	//	j, _ := json.Marshal(params)
	//	b.mqClient.Publish(topic, 0, string(j), false)
	//}
	// 再给订阅了该消息的插件发送内容
	if b.pluginNotice != nil {
		return b.pluginNotice(topic, params...)
	}
	return nil
}
