/**
 * @Author: Joey
 * @Description:
 * @Create Date: 2023/8/15 17:10
 */

package db

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func createSetting(s Setting) *gorm.Config {
	// 初始化gorm
	cfg := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   true,
		},
		SkipDefaultTransaction: s.SkipDefaultTransaction == 1,
	}
	if s.OpenLog == 1 {
		cfg.Logger = logger.Default.LogMode(logger.Info)
	}
	return cfg
}
func NewSqlite(path string, s Setting) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path), createSetting(s))
	if s.JournalMode != "" {
		db.Exec(fmt.Sprintf("PRAGMA journal_mode = %s;", s.JournalMode))
	}
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

type Setting struct {
	//DBType                 string `comment:"数据库类型：sqlite, sqlserver\n 参数\n sqlite：xxx.db\n sqlserver：用户名:密码@地址?database=数据库&encrypt=disable"`
	OpenLog                byte   `comment:"是否输出脚本日志 0否 1是"`
	SkipDefaultTransaction byte   `comment:"跳过默认事务 0否 1是"`
	JournalMode            string `comment:"Journal模式\n DELETE：在事务提交后，删除journal文件\n MEMORY：在内存中生成journal文件，不写入磁盘\n WAL：使用WAL（Write-Ahead Logging）模式，将journal记录写入WAL文件中\n OFF：完全关闭journal模式，不记录任何日志消息"`
}
