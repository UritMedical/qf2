package qdefine

import "fmt"

const (
	ErrorCodeSubmit  = "501" // 提交操作失败（增删改）
	ErrorCodeTimeOut = "502" // 执行超时
	ErrorCodeOSError = "503" // 系统故障，不可恢复
)

var errorCodeTextMap = map[string]string{
	ErrorCodeSubmit:  "执行失败",
	ErrorCodeTimeOut: "执行超时",
	ErrorCodeOSError: "系统故障",
}

const (
	RefuseCodeTokenInvalid     = "401" // Token无效或者已过期，需要重新登录
	RefuseCodePermissionDenied = "402" // 权限不足，拒绝访问
	RefuseCodeParamInvalid     = "403" // 传入参数无效
	RefuseCodeNoRecord         = "404" // 记录不存在
)

var refuseCodeTextMap = map[string]string{
	RefuseCodeTokenInvalid:     "Token无效或过期，请重新登录",
	RefuseCodePermissionDenied: "权限不足，拒绝访问",
	RefuseCodeParamInvalid:     "传入的参数无效",
	RefuseCodeNoRecord:         "记录不存在",
}

type Error struct {
	code  string
	desc  string
	error string
}

type Refuse struct {
	code string
	desc string
}

func NewError(code string, err error) Error {
	msg := ""
	if err != nil {
		msg = err.Error()
	}
	return Error{
		code:  fmt.Sprintf("%v", code),
		desc:  errorCodeTextMap[code],
		error: msg,
	}
}

func NewRefuse(code string, desc string) Refuse {
	tmp := refuseCodeTextMap[code]
	if tmp != "" {
		desc = tmp
	}
	return Refuse{
		code: fmt.Sprintf("%v", code),
		desc: desc,
	}
}

func (e Error) Code() string {
	return e.code
}

func (e Error) Desc() string {
	return e.desc
}

func (e Error) Error() string {
	return e.error
}

func (r Refuse) Code() string {
	return r.code
}

func (r Refuse) Desc() string {
	return r.desc
}

func (r Refuse) Error() string {
	return ""
}
