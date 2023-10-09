package plugin_user

import (
	"github.com/UritMedical/qf2/define"
	"github.com/UritMedical/qf2/example/plugin_user/blls/login"
	"github.com/UritMedical/qf2/example/plugin_user/blls/manage"
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
		login.New(nil),
		manage.New(nil),
	)
}

func (p *Plugin) GetId() string {
	return "User"
}

func (p *Plugin) GetVersion() string {
	return "V1.0.1.01.230901B01"
}
