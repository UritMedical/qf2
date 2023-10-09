package qhandler

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

func Exec(funcHandler interface{}, params ...interface{}) (interface{}, error) {
	funcType := reflect.TypeOf(funcHandler)
	funcParamCount := funcType.NumIn()
	if funcParamCount != len(params) {
		// 参数个数不一致
		return nil, errors.New("")
	}

	// 获取
	objects := make([]reflect.Value, funcParamCount)
	for i := 0; i < funcType.NumIn(); i++ {
		ta := funcType.In(i)
		tb := reflect.TypeOf(params[i])
		if ta == tb {
			// 直接类型转换
			objects[i] = reflect.ValueOf(params[i])
		} else {
			// 使用json转换
			obj := reflect.New(ta).Interface()
			obj2 := &obj
			j, _ := json.Marshal(params[i])
			_ = json.Unmarshal(j, obj2)
			obj3 := reflect.ValueOf(obj2).Elem().Elem().Elem().Interface()
			objects[i] = reflect.ValueOf(obj3)
			fmt.Println("")
		}
	}

	// 执行方法
	funcValue := reflect.ValueOf(funcHandler)
	returns := funcValue.Call(objects)

	// 返回结果
	rtObj := returns[0].Interface()
	rtErr := returns[1].Interface()
	if rtErr == nil {
		return rtObj, nil
	}
	return rtObj, rtErr.(error)
}

func ExecNotReturn(funcHandler interface{}, params ...interface{}) {
	funcType := reflect.TypeOf(funcHandler)
	funcParamCount := funcType.NumIn()
	if funcParamCount != len(params) {
		// 参数个数不一致
		return
	}

	// 获取
	objects := make([]reflect.Value, funcParamCount)
	for i := 0; i < funcType.NumIn(); i++ {
		ta := funcType.In(i)
		tb := reflect.TypeOf(params[i])
		if ta == tb {
			// 直接类型转换
			objects[i] = reflect.ValueOf(params[i])
		} else {
			// 使用json转换
			obj := reflect.New(ta).Interface()
			obj2 := &obj
			j, _ := json.Marshal(params[i])
			_ = json.Unmarshal(j, obj2)
			obj3 := reflect.ValueOf(obj2).Elem().Elem().Elem().Interface()
			objects[i] = reflect.ValueOf(obj3)
			fmt.Println("")
		}
	}

	// 执行方法
	funcValue := reflect.ValueOf(funcHandler)
	_ = funcValue.Call(objects)
}
