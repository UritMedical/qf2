/**
 * @Author: Joey
 * @Description:
 * @Create Date: 2023/8/12 13:46
 */

package main

import (
	"fmt"
	. "qf/define"
)

type pluginA struct {
	QBasePlugin
	name string
}

func NewPluginA(name string) QPlugin {
	p := &pluginA{}
	p.name = name
	return p
}
func (p *pluginA) Apis() (apis map[string]ApiHandler) {
	return map[string]ApiHandler{
		"Hello": p.hello,
	}
}
func (p *pluginA) Subscribes() (notices map[string]NoticeHandler) {
	return map[string]NoticeHandler{
		"topic": p.onTopic,
	}
}

func (p *pluginA) GetId() string {
	return p.name
}

func (p *pluginA) hello(params map[string]interface{}) (interface{}, error) {
	input := params["input"].(string)
	defer p.SendNotice("topic", false, p.name+" send a async Notice from hello")
	p.Logger().Debug("%s say Hello with  %s", p.name, input)
	return p.name + " say Hello with  " + input, nil
}

func (p *pluginA) onTopic(msg interface{}) {
	fmt.Printf("%s got a notice from topic \"%s\"\r\n", p.name, msg.(string))
}
