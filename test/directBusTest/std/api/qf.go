/**
 * @Author: Joey
 * @Description:
 * @Create Date: 2023/8/15 18:25
 */

package api

import (
	. "github.com/UritMedical/qf2/define"
	. "github.com/UritMedical/qf2/test/directBusTest/std/qDefine"
)

type plugin struct {
	QBasePlugin
	onTubeDelete func(tube Tube)
}

func (p *plugin) Apis() (apis map[string]ApiHandler) {
	return nil
}
func (p *plugin) Subscribes() (notices map[string]NoticeHandler) {
	return map[string]NoticeHandler{
		"TubeDelete": func(param interface{}) {
			p.onTubeDelete(param.(Tube))
		},
	}
}
func (p *plugin) GetId() string {
	return "API"
}
func (p *plugin) invokeTubeInsert(tube Tube) QError {
	args := map[string]interface{}{"model": tube}
	_, err := p.Invoke("TubeInsert", args)

	return err
}
func (p *plugin) invokeTubeDelete(id uint64) QError {
	args := map[string]interface{}{"id": id}
	_, err := p.Invoke("TubeDelete", args)
	return err
}
