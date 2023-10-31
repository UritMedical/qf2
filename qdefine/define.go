package qdefine

import (
	"reflect"
)

type QAdapter interface {
	RegApi(name string, handler QApiHandler)
	RegSubscribe(name string, handler QNoticeHandler)
	SendNotice(topic string, payload interface{})
	Invoke(appName, funcName string, header, query, body map[string]interface{}) (interface{}, QFail)
}

type QBllSvc interface {
	Init()
	Bind()
	Middleware(funcName string, ctx QContext) QFail
	Stop()
}

type QDaoSvc interface {
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
}

type QFail interface {
	Code() int
	Desc() string
	Error() string
}

type QApiHandler func(ctx QContext) (interface{}, QFail)

type QNoticeHandler func(ctx QContext)
