package manage

import (
	"github.com/UritMedical/qf2/define"
	"github.com/UritMedical/qf2/example/plugin_user/base"
	"gorm.io/gorm"
)

type bll struct {
	baseBll
}

func New(db *gorm.DB) define.QBll {
	b := &bll{
		baseBll: baseBll{},
	}
	b.funcGetUserModel = b.GetUserModel
	b.funcGetUserModelByPhone = b.GetUserModelByPhone
	b.funcUpdateUserInfo = b.UpdateUserInfo
	return b
}

func (b *bll) Stop() {

}

func (b *bll) GetUserModel(id uint) (base.UserInfo, error) {
	//face, _ := apis.Resource_GetResourceInfo(fmt.Sprintf("%d", id))

	return base.UserInfo{
		Id:    id,
		Phone: "13411112222",
		Name:  "张三",
		Sex:   "男",
		Pwd:   "fdafdasfdsafdasfdsafdsafdas",
		Face:  "",
	}, nil
}

func (b *bll) GetUserModelByPhone(phone string) (base.UserInfo, error) {
	// 假设通过数据库查询到id
	id := uint(1)

	//face, _ := base.Resource_GetResourceInfo(fmt.Sprintf("%d", id))

	return base.UserInfo{
		Id:    id,
		Phone: phone,
		Name:  "张三",
		Sex:   "男",
		Pwd:   "fdafdasfdsafdasfdsafdsafdas",
		Face:  "",
	}, nil
}

func (b *bll) UpdateUserInfo(model base.UserInfo) (bool, error) {
	// 假设通过数据库更新

	// 成功后，发送通知
	_ = b.Send_InfoChanged(model.Id)

	return false, nil
}
