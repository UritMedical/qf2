/**
 * @Author: Joey cgo包裹器将把所有的api，消息的输入输出参数以json统一序列化进行调用
 * @Description:
 * @Create Date: 2023/7/26 19:25
 */

package cgo

/*
#include <stdio.h>
typedef void (*LogGotCallBack)(char* ,int);//日志回调函数

static void Run_LogGotCallBack(LogGotCallBack cb, char* data,int len){
	cb(data,len);
}

typedef void (*NoticeGotCallBack)(char* ,int);//消息回调函数 参数以json统一序列化进行调用

static int Run_NoticeGotCallBack(NoticeGotCallBack cb, char* data,int len){
	return	cb(data,len);
}
typedef void (*ReferenceCallBack)(char* ,int,char**,int*);//引用回调函数 参数以json统一序列化进行调用

static int Run_ReferenceCallBack(ReferenceCallBack cb, char* data,int len, char** ref, int* len_ref){
	return	cb(data,len,ref,len_ref);
}

*/
import "C"
import (
	"encoding/json"
	"unsafe"
)

var (
	onLogGot    C.LogGotCallBack
	onNoticeGot C.NoticeGotCallBack
	onReference C.ReferenceCallBack

	configPath string
	bus        define.IBus
	adapter    define.IPlugin
)

//export Start
func Start() int {
	err := bus.Start()
	if err != nil {
		sendLog(err.Error())
		return 0
	}
	return 1
}

//export SendBarcode
func Invoke(paramJson *C.char, paramLength C.int, ref **C.char, refLength *C.int) C.int {
	param := &ApiParam{}
	us := unsafe.Pointer(paramJson)
	paramStr := C.GoBytes(us, paramLength)

	err := json.Unmarshal([]byte(string(paramStr)), param)
	if err != nil {
		sendLog(err.Error())
		return 0
	}
	err = adapter.Invoke(param.Route, param.Args)
	if err != nil {
		sendLog(err.Error())
		return 0
	}
	return 1
	adapter.Invoke()
	return C.int(1)

}
