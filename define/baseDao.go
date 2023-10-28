package define

import (
	"errors"
	"fmt"
	"github.com/UritMedical/qf2/utils/qreflect"
	"gorm.io/gorm"
)

type DbSimple struct {
	Id uint64 `gorm:"primaryKey"` // 唯一号
}

type DbModel struct {
	Id       uint64   `gorm:"primaryKey"` // 唯一号
	LastTime DateTime `gorm:"index"`      // 最后操作时间时间
	Summary  string   // 摘要
	FullInfo string   // 其他扩展内容
}

// Other
//
//	@Description: 其他数据
//	@return map[string]interface{}
func (model DbModel) Other() Expand {
	ref := qreflect.New(model)
	return ref.ToMapExpandAll()
}

func (model DbModel) FromModel(obj any) {
	fmt.Println(model)
}

type Expand map[string]interface{}

func (expand Expand) GetString(key string) string {
	if str, ok := expand[key].(string); ok {
		return str
	}
	return ""
}

type BaseDao[T any] struct {
	db *gorm.DB
}

func NewBaseDao[T any](db *gorm.DB) *BaseDao[T] {
	// 主动创建数据库
	err := db.AutoMigrate(new(T))
	if err != nil {
		return nil
	}
	return &BaseDao[T]{db: db}
}

// DB
//
//	@Description: 返回数据库
//	@return *gorm.DB
func (dao *BaseDao[T]) DB() *gorm.DB {
	return dao.db
}

// Create
//
//	@Description: 新建一条记录
//	@param model 对象
//	@return bool
//	@return error
func (dao *BaseDao[T]) Create(model *T) (bool, error) {
	ref := qreflect.New(model)
	if ref.Get("LastTime") == "0001-01-01 00:00:00" {
		_ = ref.Set("LastTime", NowDateTime())
	}
	// 提交
	result := dao.DB().Create(model)
	if result.RowsAffected > 0 {
		return true, nil
	}
	if result.Error != nil {
		return false, result.Error
	}
	return false, nil
}

// Update
//
//	@Description: 修改一条记录
//	@param model 对象
//	@return bool
//	@return error
func (dao *BaseDao[T]) Update(model *T) (bool, error) {
	ref := qreflect.New(model)
	if ref.Get("LastTime") == "0001-01-01 00:00:00" {
		_ = ref.Set("LastTime", NowDateTime())
	}
	// 提交
	result := dao.DB().Updates(model)
	if result.RowsAffected > 0 {
		return true, nil
	}
	if result.Error != nil {
		return false, result.Error
	}
	return false, errors.New("update record does not exist")
}

// Save
//
//	@Description: 修改一条记录（不存在则新增）
//	@param model 对象
//	@return bool
//	@return error
func (dao *BaseDao[T]) Save(model *T) (bool, error) {
	ref := qreflect.New(model)
	if ref.Get("LastTime") == "0001-01-01 00:00:00" {
		_ = ref.Set("LastTime", NowDateTime())
	}
	// 提交
	result := dao.DB().Save(model)
	if result.RowsAffected > 0 {
		return true, nil
	}
	if result.Error != nil {
		return false, result.Error
	}
	return false, nil
}

// Delete
//
//	@Description: 删除一条记录
//	@param id 记录Id
//	@return bool
//	@return error
func (dao *BaseDao[T]) Delete(id uint64) (bool, error) {
	result := dao.DB().Where("id = ?", id).Delete(new(T))
	if result.RowsAffected > 0 {
		return true, nil
	}
	if result.Error != nil {
		return false, result.Error
	}
	return false, errors.New("delete record does not exist")
}

// GetModel
//
//	@Description: 获取一条记录
//	@param id
//	@return *T
//	@return error
func (dao *BaseDao[T]) GetModel(id uint64) (*T, error) {
	// 创建空对象
	model := new(T)
	// 查询
	result := dao.DB().Where("id = ?", id).Find(model)
	// 如果异常或者未查询到任何数据
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, result.Error
	}
	return model, nil
}

// CheckExist
//
//	@Description: 验证数据是否存在
//	@param id
//	@return bool
//	@return error
func (dao *BaseDao[T]) CheckExist(id uint64) (bool, error) {
	// 创建空对象
	model := new(T)
	// 查询
	result := dao.DB().Where("id = ?", id).Find(model)
	// 如果异常或者未查询到任何数据
	if result.Error != nil || result.RowsAffected == 0 {
		return false, result.Error
	}
	return true, nil
}

// GetList
//
//	@Description: 查询一组列表
//	@param startId 起始Id
//	@param maxCount 最大数量
//	@return []*T
//	@return error
func (dao *BaseDao[T]) GetList(startId uint64, maxCount int) ([]*T, error) {
	list := make([]*T, 0)
	// 查询
	result := dao.DB().Limit(int(maxCount)).Offset(int(startId)).Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return list, nil
}

// GetAll
//
//	@Description: 返回所有列表
//	@return []*T
//	@return error
func (dao *BaseDao[T]) GetAll() ([]*T, error) {
	list := make([]*T, 0)
	// 查询
	result := dao.DB().Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return list, nil
}

// GetCondition
//
//	@Description: 条件查询一条记录
//	@param query 条件，如 id = ? 或 id IN (?) 等
//	@param args 条件参数，如 id, ids 等
//	@return []*T
//	@return error
func (dao *BaseDao[T]) GetCondition(query interface{}, args ...interface{}) (*T, error) {
	model := new(T)
	// 查询
	result := dao.DB().Where(query, args).Find(model)
	if result.Error != nil {
		return nil, result.Error
	}
	return model, nil
}

// GetConditions
//
//	@Description: 条件查询一组列表
//	@param query 条件，如 id = ? 或 id IN (?) 等
//	@param args 条件参数，如 id, ids 等
//	@return []*T
//	@return error
func (dao *BaseDao[T]) GetConditions(query interface{}, args ...interface{}) ([]*T, error) {
	list := make([]*T, 0)
	// 查询
	result := dao.DB().Where(query, args).Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return list, nil
}