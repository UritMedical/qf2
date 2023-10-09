package base

import "github.com/UritMedical/qf2/define"

var Plugin *define.BasePlugin

type LoginInfo struct {
	Id    uint     // 用户代号
	Phone string   // 用户手机号
	Token string   // 用户token
	Info  UserInfo // 用户信息
}

type UserInfo struct {
	Id    uint   // 用户代号
	Phone string // 用户手机号
	Name  string // 用户姓名
	Sex   string // 用户性别
	Pwd   string // 用户密码
	Face  string // 用户头像
}
