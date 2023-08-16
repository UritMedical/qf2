/**
 * @Author: Joey
 * @Description:
 * @Create Date: 2023/8/15 20:53
 */

package api

import (
	"fmt"
	. "github.com/UritMedical/qf2/define"
	. "github.com/UritMedical/qf2/test/directBusTest/std/qdefine"
)

type Bll struct {
	plugin
}

func New() *Bll {
	b := &Bll{
		plugin: plugin{},
	}
	b.plugin.onTubeDelete = b.onTubeDelete
	return b
}
func (b *Bll) onTubeDelete(tube Tube) {
	//todo Send MQTT msg to web front end
	fmt.Printf("got a notice from bus on tubedelete %v", tube)
}
func (b *Bll) TubeInsert(tube Tube) QError {
	return b.invokeTubeInsert(tube)
}
func (b *Bll) TubeDelete(id uint64) QError {
	return b.invokeTubeDelete(id)
}
