/**
 * @Author: Joey
 * @Description:
 * @Create Date: 2023/8/15 15:27
 */

package tube

import (
	. "github.com/UritMedical/qf2/define"
	. "github.com/UritMedical/qf2/test/directBusTest/std/qdefine"
)

type plugin struct {
	QBasePlugin
	tubeInsert func(model Tube) QError
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
	return nil
}

func (p *plugin) GetId() string {
	return "Tube"
}
func (p *plugin) SendNoticeTubeDelete(tube Tube, isSync bool) {
	p.SendNotice("TubeDelete", isSync, tube)
}
