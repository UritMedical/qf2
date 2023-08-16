/**
 * @Author:  QPluginBuilder
 * @Description:试管模块
 * @Create Date: 2023-08-16 19:23:42
 */

package tube

import (
	. "github.com/UritMedical/qf2/define"
	. "github.com/UritMedical/qf2/test/directBusTest/std/qDefine"
)

type plugin struct {
	QBasePlugin
	//创建试管
	tubeInsert func(model Tube) QError
	//删除试管
	tubeDelete func(id uint64) QError
}

func (p *plugin) Apis() (apis map[string]ApiHandler) {
	return map[string]ApiHandler{
		"TubeInsert": func(params map[string]interface{}) (interface{}, QError) {
			model := params["model"].(Tube)
			return nil, p.tubeInsert(model)
		},
		"TubeDelete": func(params map[string]interface{}) (interface{}, QError) {
			id := params["id"].(uint64)
			return nil, p.tubeDelete(id)
		},
	}
}
func (p *plugin) Subscribes() (notices map[string]NoticeHandler) {
	return map[string]NoticeHandler{}
}

func (p *plugin) GetId() string {
	return "Test.Tube"
}

// SendNoticeTubeDeleted 发送试管删除消息
func (p *plugin) SendNoticeTubeDeleted(tube Tube, isSync bool) QError {
	return p.SendNotice("TubeDeleted", isSync, tube)
}
