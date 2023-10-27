package define

import "gorm.io/gorm"

type BaseDao struct {
	db *gorm.DB
}

func NewBaseDao(db *gorm.DB) BaseDao {
	return BaseDao{db: db}
}

func (dao BaseDao) DB() *gorm.DB {
	return dao.db
}
