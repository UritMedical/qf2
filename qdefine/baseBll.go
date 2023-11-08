package qdefine

type BaseBll struct {
	Fail fail
}

type fail struct {
}

// NoRecord
//
//	@Description: 记录不存在
//	@param desc
//	@return Refuse
func (f fail) Error(err error) *QFail {
	return &QFail{
		Code: 0,
		Err:  err,
	}
}

// Refuse
//
//	@Description: 自定义拒绝
//	@param code
//	@param desc
//	@return Refuse
func (f fail) Refuse(code int, desc string) *QFail {
	return &QFail{
		Code: code,
		Err:  nil,
		Desc: desc,
	}
}
