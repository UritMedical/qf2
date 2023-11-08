package qdefine

const (
	ErrorCodeSubmit  = 501 + iota // 提交操作失败（增删改）
	ErrorCodeTimeOut              // 执行超时
	ErrorCodeOSError              // 系统故障，不可恢复

)

//var errorCodeTextMap = map[string]string{
//	ErrorCodeSubmit:  "执行失败",
//	ErrorCodeTimeOut: "执行超时",
//	ErrorCodeOSError: "系统故障",
//}

const (
	RefuseCodeTokenInvalid     = 401 + iota // Token无效或者已过期，需要重新登录
	RefuseCodePermissionDenied              // 权限不足，拒绝访问
	RefuseCodeParamInvalid                  // 传入参数无效
	RefuseCodeNoRecord                      // 记录不存在
	RefuseCodeRecordExist                   //记录已存在
	RefusePwdError                          //密码错误
)

//var refuseCodeTextMap = map[int]string{
//	RefuseCodeTokenInvalid:     "Token无效或过期，请重新登录",
//	RefuseCodePermissionDenied: "权限不足，拒绝访问",
//	RefuseCodeParamInvalid:     "传入的参数无效",
//	RefuseCodeNoRecord:         "记录不存在",
//}

//type Error struct {
//	code  string
//	desc  string
//	error string
//}

//type Refuse struct {
//	code string
//	desc string
//}

//func NewError(err error) Error {
//	msg := ""
//	if err != nil {
//		msg = err.Error()
//	}
//	return Error{
//		code:  fmt.Sprintf("%v", code),
//		desc:  errorCodeTextMap[code],
//		error: msg,
//	}
//}

//func NewRefuse(code int, ) Refuse {
//	//tmp := refuseCodeTextMap[code]
//	//if tmp != "" {
//	//	desc = tmp
//	//}
//	return Refuse{
//		code: fmt.Sprintf("%v", code),
//		desc: desc,
//	}
//}
//
//func (e Error) Code() string {
//	return e.code
//}
//
//func (e Error) Desc() string {
//	return e.desc
//}
//
//func (e Error) Error() string {
//	return e.error
//}
//
//func (r Refuse) Code() string {
//	return r.code
//}
//
//func (r Refuse) Desc() string {
//	return r.desc
//}
//
//func (r Refuse) Error() string {
//	return ""
//}
