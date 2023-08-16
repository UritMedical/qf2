/**
 * @Author: Joey
 * @Description:
 * @Create Date: 2023/8/12 13:46
 */

package main

import (
	. "github.com/UritMedical/qf2/define"
)

type pluginB struct {
	pluginA
}

func NewPluginB(name string) QPlugin {
	p := &pluginB{}
	p.name = name
	return p
}
func (p *pluginB) Apis() (apis map[string]ApiHandler) {
	return map[string]ApiHandler{
		"GoodBye": p.goodBye,
	}
}
func (p *pluginB) goodBye(params map[string]interface{}) (interface{}, QError) {
	input := params["input"].(string)
	defer p.SendNotice("topic", true, p.name+" send a sync Notice from good bye")
	return p.name + " say good bye with  " + input, nil
}
