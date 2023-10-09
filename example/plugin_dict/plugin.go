package plugin_dict

import (
	"github.com/UritMedical/qf2/define"
	"github.com/UritMedical/qf2/example/plugin_dict/blls/dict"
)

type Plugin struct {
	define.BasePlugin
}

type setting struct {
}

func (p *Plugin) DefaultSetting() interface{} {
	return setting{}
}

func (p *Plugin) Init(setting interface{}) {
	// 注册业务
	p.RegBll(
		dict.New(nil),
	)
}

func (p *Plugin) GetId() string {
	return "Dict"
}

func (p *Plugin) GetVersion() string {
	return "V1.0.1.01.230901B01"
}
