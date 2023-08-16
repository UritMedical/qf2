/**
 * @Author: Joey
 * @Description:
 * @Create Date: 2023/8/16 9:14
 */

package define

import (
	"fmt"
	"github.com/UritMedical/qf2/util/db"
	"github.com/UritMedical/qf2/util/qreflect"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

//
// BaseDao
//  @Description: 基础数据访问方法
//
type BaseDao struct {
	db        *gorm.DB
	tableName string
}

func NewBaseDao(db *gorm.DB, model interface{}, migrator bool) BaseDao {
	dao := BaseDao{}
	dao.init(db, model, migrator)
	return dao
}

func NewSqlite(path string) *gorm.DB {
	return db.NewSqlite(path, db.Setting{
		OpenLog:                0,
		SkipDefaultTransaction: 1,
		JournalMode:            "OFF",
	})
}

//
// init
//  @Description: 初始化数据库
//  @param db
//  @param model
//  @param migrator
//
func (b *BaseDao) init(db *gorm.DB, model interface{}, migrator bool) {
	b.db = db
	// 根据实体名称，生成数据库
	if model != nil {
		b.tableName = buildTableName(model)
		// 自动生成表
		if migrator {
			// 每次都生成
			err := db.Table(b.tableName).AutoMigrate(model)
			if err != nil {
				panic(fmt.Sprintf("AutoMigrate %s failed: %s", b.tableName, err.Error()))
			}
		} else {
			// 仅第一次生成
			if db.Migrator().HasTable(b.tableName) == false {
				err := db.Table(b.tableName).AutoMigrate(model)
				if err != nil {
					panic(fmt.Sprintf("AutoMigrate %s failed: %s", b.tableName, err.Error()))
				}
			}
		}
	}
}

//
// buildTableName
//  @Description: 根据结构体，生成对应的数据库表名
//  @param model 结构体
//  @return string 然后表名，规则：包名_结构体名，如果包名和结构体名一致时，则只返回结构体名
//
func buildTableName(model interface{}) string {
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	per := ""
	// 如果是框架内部业务，则直接增加Qf前缀
	// 反之直接使用实体名称
	if strings.HasPrefix(strings.ToLower(t.PkgPath()), "github.com/uritmedical/qf") {
		per = "Qf"
	}
	return fmt.Sprintf("%s%s", per, t.Name())
}

//
// DB
//  @Description: 返回对应表的数据控制器
//  @return *gorm.DB
//
func (b *BaseDao) DB() *gorm.DB {
	if b.tableName == "" {
		return b.db
	}
	return b.db.Table(b.tableName)
}

//
// Create
//  @Description: 新增内容
//  @param content 包含了内容结构的实体对象
//  @return IError 异常
//
func (b *BaseDao) Create(content interface{}) error {
	ref := qreflect.New(content)
	if ref.Get("LastTime") == "0001-01-01 00:00:00" {
		_ = ref.Set("LastTime", NowDateTime())
	}
	// 提交
	result := b.DB().Create(content)
	if result.RowsAffected > 0 {
		return nil
	}
	if result.Error != nil {
		return Error(ErrorCodeSaveFailure, result.Error.Error())
	}
	return nil
}

//
// Save
//  @Description: 保存内容（存在更新、不存在则新增）
//  @param content 包含了内容结构的实体对象
//  @return IError 异常
//
func (b *BaseDao) Save(content interface{}) QError {
	ref := qreflect.New(content)
	if ref.Get("LastTime") == "0001-01-01 00:00:00" {
		_ = ref.Set("LastTime", NowDateTime())
	}
	// 提交
	result := b.DB().Save(content)
	if result.RowsAffected > 0 {
		return nil
	}
	if result.Error != nil {
		return Error(ErrorCodeSaveFailure, result.Error.Error())
	}
	return nil
}

//
// Delete
//  @Description: 删除内容
//  @param id 唯一号
//  @return QError 异常
//
func (b *BaseDao) Delete(id uint64) QError {
	result := b.DB().Delete(&BaseModel{Id: id})
	if result.RowsAffected == 0 {
		return Error(ErrorCodeRecordNotFound, fmt.Sprintf("delete failed, id=%d does not exist", id))
	}
	if result.Error != nil {
		return Error(ErrorCodeDeleteFailure, result.Error.Error())
	}
	return nil
}

//
// GetModel
//  @Description: 获取单条数据
//  @param id 唯一号
//  @param dest 目标实体结构
//  @return QError 返回异常
//
func (b *BaseDao) GetModel(id uint64, dest interface{}) QError {
	result := b.DB().Where("Id = ?", id).Find(dest)
	// 如果异常或者未查询到任何数据
	if result.Error != nil {
		return Error(ErrorCodeRecordNotFound, result.Error.Error())
	}
	return nil
}

//
// GetSummary
//  @Description: 仅获取单条摘要
//  @param id
//  @return QError
//
func (b *BaseDao) GetSummary(id uint64, dest interface{}) QError {
	result := b.DB().Where("Id = ?", id).Omit("FullInfo").Find(dest)
	// 如果异常或者未查询到任何数据
	if result.Error != nil {
		return Error(ErrorCodeRecordNotFound, result.Error.Error())
	}
	return nil
}

//
// GetList
//  @Description: 按唯一号区间，获取一组列表
//  @param startId 起始编号
//  @param maxCount 最大获取数
//  @param dest 目标列表
//  @return QError 返回异常
//
func (b *BaseDao) GetList(startId uint64, maxCount uint, dest interface{}) QError {
	result := b.DB().Limit(int(maxCount)).Offset(int(startId)).Find(dest)
	if result.Error != nil {
		return Error(ErrorCodeRecordNotFound, result.Error.Error())
	}
	return nil
}

//
// GetListByIN
//  @Description: 获取指定Id列表的一组数据列表
//  @param ids Id列表
//  @param dest
//  @return QError
//
func (b *BaseDao) GetListByIN(ids []uint64, dest interface{}) QError {
	result := b.DB().Where("Id IN (?)", ids).Find(dest)
	if result.Error != nil {
		return Error(ErrorCodeRecordNotFound, result.Error.Error())
	}
	return nil
}

//
// GetConditions
//  @Description: 通过自定义条件获取数据
//  @param dest 结构体/列表
//  @param query 条件
//  @param args 条件参数
//  @return QError
//
func (b *BaseDao) GetConditions(dest interface{}, query interface{}, args ...interface{}) QError {
	result := b.DB().Where(query, args...).Find(dest)
	if result.Error != nil {
		return Error(ErrorCodeRecordNotFound, result.Error.Error())
	}
	return nil
}

//
// GetCount
//  @Description: GetCount
//  @param query 查询条件，如：a = ? and b = ?
//  @param args 条件对应的值
//  @return int64 查询到的记录数
//
func (b *BaseDao) GetCount(query interface{}, args ...interface{}) int64 {
	count := int64(0)
	b.DB().Where(query, args).Count(&count)
	return count
}

//
// CheckExists
//  @Description: 检查内容是否存在
//  @param id 唯一号
//  @return bool true存在 false不存在
//
func (b *BaseDao) CheckExists(id uint64) bool {
	count := int64(0)
	result := b.DB().Where("Id = ?", id).Count(&count)
	if count > 0 && result.Error == nil {
		return true
	}
	return false
}
