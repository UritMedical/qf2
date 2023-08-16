/**
 * @Author: Joey
 * @Description:
 * @Create Date: 2023/8/15 15:40
 */

package result

import (
	. "github.com/UritMedical/qf2/define"
	. "github.com/UritMedical/qf2/test/directBusTest/std/qDefine"
	"gorm.io/gorm"
)

type bll struct {
	dao dao
	plugin
}

func New(db *gorm.DB) QPlugin {
	bll := &bll{
		plugin: plugin{},
		dao: dao{
			BaseDao: NewBaseDao(db, Result{}, false),
		},
	}
	bll.plugin.resultInsert = bll.resultInsert
	bll.plugin.resultDelete = bll.resultDelete
	bll.plugin.onTubeDeleted = bll.onTubeDeleted
	return bll
}
func (b *bll) resultInsert(model Result) QError {
	return b.dao.Save(&model)
}
func (b *bll) resultDelete(id uint64) QError {
	return b.dao.Delete(id)
}

func (b *bll) onTubeDeleted(tube Tube) {
	e := b.dao.deleteWithTube(tube.Id)
	if e != nil {
		b.Logger().Error("%v", e)
	}
}
