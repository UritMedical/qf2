/**
 * @Author:  QPluginBuilder
 * @Description:检验结果模块
 * @Create Date: 2023-08-16 19:49:18
 */

package result

import (
	. "github.com/UritMedical/qf2/define"
	. "github.com/UritMedical/qf2/test/directBusTest/std/qDefine"
)

type plugin struct {
	QBasePlugin
	//创建检验结果
	resultInsert func(model Result) QError
	//删除检验结果
	resultDelete func(id uint64) QError

	//试管删除
	onTubeDeleted func(tube Tube)
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
		"TubeDeleted": func(param interface{}) {
			p.onTubeDeleted(param.(Tube))
		},
	}
}

func (p *plugin) GetId() string {
	return "Test.Result"
}
