package base

func Login(phone string, password string) (LoginInfo, error) {
	rs, err := Plugin.Invoke("Login", phone, password)
	return rs.(LoginInfo), err
}

func CheckToken(token string) (bool, error) {
	rs, err := Plugin.Invoke("CheckToken", token)
	return rs.(bool), err
}

func GetUserModel(id uint) (UserInfo, error) {
	rs, err := Plugin.Invoke("GetUserModel", id)
	return rs.(UserInfo), err
}

func GetUserModelByPhone(phone string) (UserInfo, error) {
	rs, err := Plugin.Invoke("GetUserModelByPhone", phone)
	return rs.(UserInfo), err
}

func UpdateUserInfo(model UserInfo) (bool, error) {
	rs, err := Plugin.Invoke("UpdateUserInfo", model)
	return rs.(bool), err
}
