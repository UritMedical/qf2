package dict

import (
	"github.com/UritMedical/qf2/define"
	"gorm.io/gorm"
)

type bll struct {
	baseBll
}

func New(db *gorm.DB) define.QBll {
	b := &bll{
		baseBll: baseBll{},
	}
	return b
}

func (b *bll) Stop() {

}
