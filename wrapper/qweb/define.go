package qweb

import (
	"github.com/UritMedical/qf2/qdefine"
	"github.com/UritMedical/qf2/utils/qconfig"
	"github.com/UritMedical/qf2/utils/qdb"
	"github.com/UritMedical/qf2/utils/qio"
)

type Module struct {
	//RouteGroup string
	//QAdapter func(adapter qdefine.QAdapter)
	QBll qdefine.QBll
	QDao qdefine.QDao
}

type setting struct {
	configPath string
	Port       int                    `comment:"服务端口"`
	DefGroup   string                 `comment:"默认路由组"`
	GormConfig map[string]qdb.Setting `comment:"gorm配置"`
	//CallConfig  callConfig             `comment:"插件间api访问地址，如果是通过启动器统一启动，则无需配置"`
	OtherConfig otherConfig `comment:"其他配置"`
}

//type callConfig struct {
//	MqClient string
//	ApiUrls  map[string]string
//}

type otherConfig struct {
	JsonDateFormat string `comment:"框架日期的Json格式"`
	JsonTimeFormat string `comment:"框架时间的Json格式"`
}

func newSetting(configPath string) *setting {
	s := &setting{
		Port:     10001,
		DefGroup: "api",
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

//type route struct {
//	Url  string
//	Type string
//}
