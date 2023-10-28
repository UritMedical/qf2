package web

import (
	"github.com/UritMedical/qf2/define"
	"github.com/UritMedical/qf2/utils/qconfig"
	"github.com/UritMedical/qf2/utils/qio"
)

type WebParams struct {
	AppInitFunc []func(adapter define.QAdapter)
	Bll         define.QSvc
}

func LoadParams(bll define.QSvc, initFunc ...func(adapter define.QAdapter)) WebParams {
	setting := WebParams{
		AppInitFunc: initFunc,
		Bll:         bll,
	}
	return setting
}

type setting struct {
	Port        int
	DefGroup    string
	MqBroker    mqConfig    `comment:"mq配置"`
	ApiNotice   apiNotice   `comment:"api通知配置"`
	OtherConfig otherConfig `comment:"其他配置"`
}

type mqConfig struct {
	Port   int `comment:"MqBroker的端口"`
	WsPort int `comment:"WebSocket的端口"`
}

type apiNotice struct {
	MqServer string   `comment:"MqBroker的地址"`
	Exposes  []string `comment:"需要发送的Api类型，如post, put, delete等"`
}

type otherConfig struct {
	JsonDateFormat string `comment:"框架日期的Json格式"`
	JsonTimeFormat string `comment:"框架时间的Json格式"`
}

func loadSetting() *setting {
	setting := &setting{
		Port:     10001,
		DefGroup: "",
	}
	setting.OtherConfig = otherConfig{
		JsonDateFormat: "yyyy-MM-dd",
		JsonTimeFormat: "HH:mm:ss",
	}
	// 加载配置文件
	path := qio.GetFullPath("./config/config.toml")
	_ = qconfig.LoadFromToml(path, setting)
	return setting
}

type route struct {
	Url  string
	Type string
}
