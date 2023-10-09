package manage

import (
	"github.com/UritMedical/qf2/define"
	"github.com/UritMedical/qf2/example/plugin_user/base"
	"github.com/UritMedical/qf2/utils/qhandler"
)

type baseBll struct {
	define.BaseBll
	funcGetUserModel        func(id uint) (base.UserInfo, error)
	funcGetUserModelByPhone func(phone string) (base.UserInfo, error)
	funcUpdateUserInfo      func(model base.UserInfo) (bool, error)
}

func (b *baseBll) GetConfig() string {
	return ""
}

func (b *baseBll) Apis() map[string]define.ApiHandler {
	return map[string]define.ApiHandler{
		"GetUserModel": func(params ...interface{}) (interface{}, error) {
			return qhandler.Exec(b.funcGetUserModel, params...)
		},
		"GetUserModelByPhone": func(params ...interface{}) (interface{}, error) {
			return qhandler.Exec(b.funcGetUserModelByPhone, params...)
		},
		"UpdateUserInfo": func(params ...interface{}) (interface{}, error) {
			return qhandler.Exec(b.funcUpdateUserInfo, params...)
		},
	}
}

func (b *baseBll) Subscribes() map[string]define.NoticeHandler {
	return nil
}

func (b *baseBll) Send_InfoChanged(uid uint) error {
	return b.SendNotice("UserUpdatedFinished", uid)
}
