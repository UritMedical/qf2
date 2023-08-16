/**
 * @Author: Joey
 * @Description:
 * @Create Date: 2023/8/15 15:27
 */

package result

import (
	. "github.com/UritMedical/qf2/define"
	. "github.com/UritMedical/qf2/test/directBusTest/std/qdefine"
)

type plugin struct {
	QBasePlugin
	resultInsert func(model Result) QError
	resultDelete func(id uint64) QError
	onTubeDelete func(tube Tube)
}

func (p *plugin) Apis() (apis map[string]ApiHandler) {
	return map[string]ApiHandler{
		"ResultInsert": func(params map[string]interface{}) (interface{}, QError) {
			model := params["model"].(Result)
			return nil, p.resultInsert(model)

		},
		"ResultDelete": func(params map[string]interface{}) (interface{}, QError) {
			id := params["id"].(uint64)
			return nil, p.resultDelete(id)
		},
	}
}
func (p *plugin) Subscribes() (notices map[string]NoticeHandler) {
	return map[string]NoticeHandler{
		"TubeDelete": func(param interface{}) {
			p.onTubeDelete(param.(Tube))
		},
	}
}
func (p *plugin) GetId() string {
	return "Result"
}
