package qdb

import (
	"errors"
	"fmt"
	"github.com/UritMedical/qf2/utils/qconfig"
	"github.com/UritMedical/qf2/utils/qio"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strings"
)

var (
	ConfigPath = "./config/config.toml"
	Settings   map[string]Setting
)

type Setting struct {
	OpenLog            byte `comment:"开关 0否 1是"`
	SkipDefTransaction byte
	DBConfig           string `comment:"数据库类型|参数\n sqlite|./db/data.db&OFF\n sqlserver|用户名:密码@地址?database=数据库&encrypt=disable\n mysql|用户名:密码@tcp(127.0.0.1:3306)/数据库?charset=utf8mb4&parseTime=True&loc=Local"`
}

// NewDb
//
//	@Description: 创建数据库
//	@param cfgSection 配置节点，用于启动多个数据库不用配置
//	@return *gorm.DB
func NewDb(cfgSection string) *gorm.DB {
	if Settings == nil {
		Settings = map[string]Setting{}
	}
	setting := loadSetting(cfgSection)
	gc := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   true,
		},
		SkipDefaultTransaction: setting.SkipDefTransaction == 1,
	}
	if setting.OpenLog == 1 {
		gc.Logger = logger.Default.LogMode(logger.Info)
	}
	sp := strings.Split(setting.DBConfig, "|")

	var db *gorm.DB
	var err error
	switch sp[0] {
	case "sqlite":
		spp := strings.Split(sp[1], "&")
		// 创建数据库
		file := qio.GetFullPath(spp[0])
		if _, err := qio.CreateDirectory(file); err != nil {
			panic(err)
		}
		db, err = gorm.Open(sqlite.Open(file), &gc)
		if err != nil {
			panic(err)
		}
		// Journal模式
		//  DELETE：在事务提交后，删除journal文件
		//  MEMORY：在内存中生成journal文件，不写入磁盘
		//  WAL：使用WAL（Write-Ahead Logging）模式，将journal记录写入WAL文件中
		//  OFF：完全关闭journal模式，不记录任何日志消息
		if spp[1] != "" {
			db.Exec(fmt.Sprintf("PRAGMA journal_mode = %s;", spp[1]))
		}
	case "sqlserver":
		dsn := fmt.Sprintf("sqlserver://%s", sp[1])
		db, err = gorm.Open(sqlserver.Open(dsn), &gc)
		if err != nil {
			panic(err)
		}
	case "mysql":
		dsn := sp[1]
		db, err = gorm.Open(mysql.Open(dsn), &gc)
		if err != nil {
			panic(err)
		}
	}
	if db == nil {
		panic(errors.New("unknown db type"))
	}
	return db
}

func loadSetting(cfgSection string) Setting {
	// 加载配置
	old := struct {
		GormConfig map[string]Setting
	}{
		GormConfig: map[string]Setting{},
	}
	_ = qconfig.OnlyLoadFromToml(qio.GetFullPath(ConfigPath), &old)
	for k, v := range old.GormConfig {
		Settings[k] = v
	}
	if _, ok := Settings[cfgSection]; ok == false {
		dbName := cfgSection
		if dbName == "" {
			dbName = "data"
		}
		Settings[cfgSection] = Setting{
			OpenLog:            0,
			SkipDefTransaction: 1,
			DBConfig:           fmt.Sprintf("sqlite|./db/%s.db&OFF", dbName),
		}
	}
	return Settings[cfgSection]
}
