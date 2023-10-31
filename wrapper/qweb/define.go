package qweb

import (
	"github.com/UritMedical/qf2/qdefine"
	"github.com/UritMedical/qf2/utils/qconfig"
	"github.com/UritMedical/qf2/utils/qdb"
	"github.com/UritMedical/qf2/utils/qio"
)

type StartParam struct {
	ConfigPath string
	QfSvc      func(adapter qdefine.QAdapter)
	BllSvc     qdefine.QBllSvc
	DaoSvc     qdefine.QDaoSvc
	Stop       func()
}

type setting struct {
	configPath  string
	Port        int                    `comment:"服务端口"`
	DefGroup    string                 `comment:"默认路由组"`
	MqClient    string                 `comment:"MQ服务地址"`
	GormConfig  map[string]qdb.Setting `comment:"gorm配置"`
	OtherConfig otherConfig            `comment:"其他配置"`
}

type otherConfig struct {
	JsonDateFormat string `comment:"框架日期的Json格式"`
	JsonTimeFormat string `comment:"框架时间的Json格式"`
}

func newSetting(configPath string) *setting {
	s := &setting{
		Port:     10001,
		DefGroup: "",
	}
	s.OtherConfig = otherConfig{
		JsonDateFormat: "yyyy-MM-dd",
		JsonTimeFormat: "HH:mm:ss",
	}
	// 加载配置文件
	_ = qconfig.OnlyLoadFromToml(qio.GetFullPath(configPath), s)
	s.configPath = configPath
	return s
}

func (s *setting) Save() {
	_ = qconfig.SaveFromToml(s.configPath, s)
}

type route struct {
	Url  string
	Type string
}
