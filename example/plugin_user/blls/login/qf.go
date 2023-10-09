package login

import (
	"github.com/UritMedical/qf2/define"
	"github.com/UritMedical/qf2/example/plugin_user/base"
	"github.com/UritMedical/qf2/utils/qhandler"
)

type baseBll struct {
	define.BaseBll
	funcLogin                    func(phone string, password string) (*base.LoginInfo, error)
	funcCheckToken               func(token string) (bool, error)
	subscribeUserUpdatedFinished func(uid uint)
}

func (b *baseBll) GetConfig() string {
	return ""
}

func (b *baseBll) Apis() map[string]define.ApiHandler {
	return map[string]define.ApiHandler{
		"Login": func(params ...interface{}) (interface{}, error) {
			return qhandler.Exec(b.funcLogin, params...)
		},
		"CheckToken": func(params ...interface{}) (interface{}, error) {
			return qhandler.Exec(b.funcCheckToken, params...)
		},
	}
}

func (b *baseBll) Subscribes() map[string]define.NoticeHandler {
	return map[string]define.NoticeHandler{
		"UserUpdatedFinished": func(params ...interface{}) {
			qhandler.ExecNotReturn(b.subscribeUserUpdatedFinished, params...)
		},
	}
}
