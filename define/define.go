package define

import (
	"reflect"
)

type QAppAdapter interface {
	RegApi(name string, handler QApiHandler)
	RegSubscribe(name string, handler QNoticeHandler)
	SendNotice(topic string, payload interface{})
	Invoke(name string, params []interface{}, of reflect.Type) interface{}
}

type QBll interface {
	Init()
	Bind()
	Stop()
}

type QContext interface {
	GetString(key string) string
	GetInt(key string) int
	GetUInt(key string) uint64
	GetBool(key string) bool
	GetAny(key string) any
}

type QFail interface {
	Code() int
	Desc() string
	Error() string
}

type QApiHandler func(ctx QContext) (interface{}, QFail)

type QNoticeHandler func(ctx QContext)
