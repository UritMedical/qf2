package qdb

import (
	"encoding/json"
	"github.com/UritMedical/qf2/utils/qreflect"
	"github.com/pkg/errors"
	"strings"
)

func Marshal[T any](apiModel any) T {
	// 先将apiModel转为字典
	js, _ := json.Marshal(apiModel)
	values := map[string]interface{}{}
	_ = json.Unmarshal(js, &values)
	// 在写入到数据库Model
	dbModel := new(T)
	_ = setModel(dbModel, values)
	return *dbModel
}

func UnMarshal[T any](dbModel any) T {
	apiModel := new(T)
	api := qreflect.New(apiModel)
	// 将数据库Model中的内容写入到apiModel中
	_ = api.SetAny(qreflect.New(dbModel).ToMapExpandAll())
	return *apiModel
}

// SetModel
//
//	@Description: 修改结构体内的字段值
//	@param objectPtr
//	@param value
func setModel(objectPtr interface{}, value map[string]interface{}) error {
	if objectPtr == nil {
		return errors.New("the object cannot be empty")
	}
	ref := qreflect.New(objectPtr)
	// 必须为指针
	if ref.IsPtr() == false {
		return errors.New("the object must be pointer")
	}

	// 修改外部值
	if value != nil {
		e := ref.SetAny(value)
		if e != nil {
			return e
		}
	}
	// 修改Info
	return setInfo(ref, value)
}

func setInfo(ref *qreflect.Reflect, value map[string]interface{}) error {
	all := ref.ToMap()

	// 复制一份
	temp := map[string]interface{}{}
	for k, v := range value {
		temp[k] = v
	}

	// 转摘要
	if field, ok := temp["SummaryFields"]; ok && field != "" {
		e := ref.Set("Summary", fields(field, all["Summary"], all, &temp))
		if e != nil {
			return e
		}
	}
	// 转信息
	if field, ok := temp["InfoFields"]; ok && field != "" {
		e := ref.Set("FullInfo", fields(field, all["FullInfo"], all, &temp))
		if e != nil {
			return e
		}
		return nil
	}

	// 将剩余的全部写入到Info中
	if info, ok := all["FullInfo"]; ok {
		mp := map[string]interface{}{}
		_ = json.Unmarshal([]byte(info.(string)), &mp)
		for k, v := range temp {
			if k == "SummaryFields" || k == "InfoFields" {
				continue
			}
			if _, ok := all[k]; ok == false {
				mp[k] = v
			}
		}
		mj, _ := json.Marshal(mp)
		e := ref.Set("FullInfo", string(mj))
		if e != nil {
			return e
		}
	}
	return nil
}

func fields(field interface{}, source interface{}, all map[string]interface{}, values *map[string]interface{}) string {
	if field == nil || field.(string) == "" {
		return ""
	}
	// 获取原始数据并转为字典
	mp := map[string]interface{}{}
	if source != nil {
		_ = json.Unmarshal([]byte(source.(string)), &mp)
	}
	// 获取需要的值
	temp := *values
	for _, name := range strings.Split(field.(string), ",") {
		if _, ok := all[name]; ok == false {
			if _, ok2 := temp[name]; ok2 {
				mp[name] = temp[name]
				delete(temp, name)
			}
		}
	}
	values = &temp
	// 返回
	if len(mp) == 0 {
		return ""
	}
	mj, _ := json.Marshal(mp)
	return string(mj)
}
