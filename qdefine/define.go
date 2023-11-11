package qdefine

import (
	"reflect"
)

// QWidget 微件
type QWidget struct {
	ConfigSection string
	Modules       map[string]QModule
}
type QModule interface {
	OnStop()
	Reg(routeGroup string, adapter QAdapter)
	OnStartInvoke(funcName string, ctx QContext) *QFail
	OnEndInvoke(funcName string, ctx QContext)
}

type QAdapter interface {
	RegGet(router string, handler QApiHandler)
	RegPost(router string, handler QApiHandler)
	RegPut(router string, handler QApiHandler)
	RegDel(router string, handler QApiHandler)

	SendNotice(topic string, payload interface{})
}

type QBll interface {
	Init()
	Bind(group string, adapter QAdapter)
	StartInvoke(funcName string, ctx QContext) QFail
	EndInvoke(funcName string, ctx QContext)
	Stop()
}

type QDao interface {
	Init()
	Stop()
}

type QContext interface {
	GetString(key string) string
	GetInt(key string) int
	GetUInt(key string) uint64
	GetByte(key string) byte
	GetBool(key string) bool
	GetDate(key string) Date
	GetTime(key string) DateTime
	GetStruct(key string, objType reflect.Type) any
	GetReturnValue() interface{}
	SetNewReturnValue(newValue interface{})
}

type QFail struct {
	Code int
	Err  error
	Desc string
}

type QApiHandler func(ctx QContext) (interface{}, *QFail)

type QNoticeHandler func(ctx QContext)

// File
// @Description: 文件
type File struct {
	Name string // 文件名
	Size int64  // 文件大小
	Data []byte // 内容
}
