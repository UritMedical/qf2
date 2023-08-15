/**
 * @Author: Joey
 * @Description:
 * @Create Date: 2023/8/12 13:46
 */

package main

import "qf/define"

type pluginB struct {
	pluginA
}

func NewPluginB(name string) define.QPlugin {
	p := &pluginB{}
	p.name = name
	return p
}
func (p *pluginB) Apis() (apis map[string]define.ApiHandler) {
	return map[string]define.ApiHandler{
		"GoodBye": p.goodBye,
	}
}
func (p *pluginB) goodBye(params map[string]interface{}) (interface{}, error) {
	input := params["input"].(string)
	defer p.SendNotice("topic", true, p.name+" send a sync Notice from good bye")
	return p.name + " say good bye with  " + input, nil
}
