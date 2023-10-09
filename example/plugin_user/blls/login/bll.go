package login

import (
	"errors"
	"fmt"
	"github.com/UritMedical/qf2/define"
	"github.com/UritMedical/qf2/example/plugin_user/base"
	"github.com/UritMedical/qf2/utils/qdate"
	"gorm.io/gorm"
	"time"
)

type bll struct {
	baseBll
	tokenCaches map[string]base.LoginInfo
}

func New(db *gorm.DB) define.QBll {
	b := &bll{
		baseBll:     baseBll{},
		tokenCaches: map[string]base.LoginInfo{},
	}
	// api方法
	b.funcLogin = b.Login
	b.funcCheckToken = b.CheckToken
	// 订阅方法
	b.subscribeUserUpdatedFinished = b.UserUpdatedFinished
	return b
}

func (b *bll) Stop() {

}

func (b *bll) Login(phone string, password string) (*base.LoginInfo, error) {
	// 获取用户
	info, _ := base.GetUserModelByPhone(phone)
	if info.Id == 0 {
		return nil, errors.New("用户不存在")
	}

	// 验证密码是否正确
	if info.Pwd != password {
		return nil, errors.New("密码错误")
	}

	login := base.LoginInfo{}
	login.Phone = phone
	login.Token = qdate.ToString(time.Now(), "yyyyMMddHHmmssfff")
	login.Info = info
	b.tokenCaches[login.Token] = login
	return &login, nil
}

func (b *bll) CheckToken(token string) (bool, error) {
	if _, ok := b.tokenCaches[token]; ok {
		return true, nil
	}
	return false, nil
}

func (b *bll) UserUpdatedFinished(uid uint) {
	fmt.Println(fmt.Sprintf("已经接收到UserUpdatedFinished消息，id=%d", uid))
}
