/**
 * @Author: Joey
 * @Description:
 * @Create Date: 2023/8/15 15:40
 */

package tube

import (
	. "github.com/UritMedical/qf2/define"
	. "github.com/UritMedical/qf2/test/directBusTest/std/qdefine"
	"gorm.io/gorm"
)

type bll struct {
	dao *dao
	plugin
}

func New(db *gorm.DB) QPlugin {
	bll := &bll{
		dao: &dao{
			BaseDao: NewBaseDao(db, Tube{}, false),
		},
		plugin: plugin{},
	}
	bll.plugin.tubeInsert = bll.tubeInsert
	bll.plugin.tubeDelete = bll.tubeDelete
	return bll
}

func (b *bll) tubeInsert(model Tube) QError {
	return b.dao.Save(&model)
}
func (b *bll) tubeDelete(id uint64) QError {
	tube := Tube{}
	r := b.dao.GetModel(id, &tube)
	if r != nil {
		return r
	}
	r = b.dao.Delete(id)
	if r == nil {
		b.SendNoticeTubeDelete(tube, true)
	}
	return r
}
