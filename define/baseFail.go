package define

const (
	ErrorCodeSubmit  = iota + 500 // 提交操作失败（增删改）
	ErrorCodeTimeOut              // 执行超时
	ErrorCodeOSError              // 系统故障，不可恢复
)

var errorCodeTextMap = map[int]string{
	ErrorCodeSubmit:  "提交操作失败",
	ErrorCodeTimeOut: "执行超时",
	ErrorCodeOSError: "系统运行故障",
}

const (
	RefuseCodeParamInvalid     = iota + 400 // 传入参数无效
	RefuseCodeTokenInvalid                  // token无效或过期
	RefuseCodePermissionDenied              // 权限不足，拒绝访问
)

var refuseCodeTextMap = map[int]string{
	RefuseCodeParamInvalid:     "传入的参数无效",
	RefuseCodeTokenInvalid:     "Token无效或过期",
	RefuseCodePermissionDenied: "权限不足，拒绝访问",
}

type Error struct {
	code  int
	desc  string
	error string
}

type Refuse struct {
	code int
	desc string
}

func NewError(code int, err error) Error {
	msg := ""
	if err != nil {
		msg = err.Error()
	}
	return Error{
		code:  code,
		desc:  errorCodeTextMap[code],
		error: msg,
	}
}

func NewRefuse(code int, desc string) Refuse {
	tmp := refuseCodeTextMap[code]
	if tmp != "" {
		desc = tmp
	}
	return Refuse{
		code: code,
		desc: desc,
	}
}

func (e Error) Code() int {
	return e.code
}

func (e Error) Desc() string {
	return e.desc
}

func (e Error) Error() string {
	return e.error
}

func (r Refuse) Code() int {
	return r.code
}

func (r Refuse) Desc() string {
	return r.desc
}

func (r Refuse) Error() string {
	return ""
}
