package qdefine

type BaseBll struct {
	Fail fail
}

type fail struct {
}

// SubmitError
//
//	@Description: 提交类故障
//	@param err
//	@return Error
func (f fail) SubmitError(err error) Error {
	return NewError(ErrorCodeSubmit, err)
}

// ExecTimeOut
//
//	@Description: 执行超时
//	@param err
//	@return Error
func (f fail) ExecTimeOut(err error) Error {
	return NewError(ErrorCodeTimeOut, err)
}

// OSError
//
//	@Description: 系统故障
//	@param err
//	@return Error
func (f fail) OSError(err error) Error {
	return NewError(ErrorCodeOSError, err)
}

// ParamInvalid
//
//	@Description: 参数无效
//	@param desc
//	@return Refuse
func (f fail) ParamInvalid(desc string) Refuse {
	return NewRefuse(RefuseCodeParamInvalid, desc)
}

// NoRecord
//
//	@Description: 记录不存在
//	@param desc
//	@return Refuse
func (f fail) NoRecord(desc string) Refuse {
	return NewRefuse(RefuseCodeNoRecord, desc)
}

// Refuse
//
//	@Description: 自定义拒绝
//	@param code
//	@param desc
//	@return Refuse
func (f fail) Refuse(code string, desc string) Refuse {
	return NewRefuse(code, desc)
}
