package qdefine

import (
	"reflect"
)

type QAdapter interface {
	RegApi(name string, handler QApiHandler)
	RegSubscribe(name string, handler QNoticeHandler)
	SendNotice(topic string, payload interface{})
	Invoke(name string, params []interface{}, of reflect.Type) interface{}
}

type QBllSvc interface {
	Init()
	Bind()
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
