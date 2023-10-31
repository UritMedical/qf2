package qdefine

import "encoding/json"

// Convert
//
//	@Description: 转换到指定对象
//	@param sourceObj
//	@return T
func Convert[T any](sourceObj interface{}) T {
	destObj := new(T)
	// 使用json序列化及反序列化
	jsonBytes, _ := json.Marshal(sourceObj)
	_ = json.Unmarshal(jsonBytes, destObj)
	return *destObj
}
