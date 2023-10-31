package qweb

import (
	"encoding/json"
	"fmt"
	"github.com/UritMedical/qf2/qdefine"
	"github.com/gin-gonic/gin"
	"github.com/gobeam/stringy"
	"io"
	"reflect"
	"strconv"
	"strings"
)

type context struct {
	gin    *gin.Context
	values *values
}

func newContext(ginCtx *gin.Context) *context {
	ctx := &context{
		gin: ginCtx,
		values: &values{
			Inputs: make([]map[string]interface{}, 0),
			//Header: map[string]interface{}{},
		},
	}
	// 加载gin的上下文中的数据
	ctx.loadValues()
	return ctx
}

func (c *context) GetString(key string) string {
	value := c.values.getValue(key)
	// 返回
	if value == nil {
		return ""
	}
	return fmt.Sprintf("%v", value)
}

func (c *context) GetInt(key string) int {
	num, _ := strconv.Atoi(c.GetString(key))
	return num
}

func (c *context) GetUInt(key string) uint64 {
	num, _ := strconv.ParseUint(c.GetString(key), 10, 64)
	return num
}

func (c *context) GetByte(key string) byte {
	num, _ := strconv.ParseInt(c.GetString(key), 10, 8)
	return byte(num)
}

func (c *context) GetBool(key string) bool {
	value := strings.ToLower(c.GetString(key))
	if value == "true" || value == "1" {
		return true
	}
	return false
}

func (c *context) GetDate(key string) qdefine.Date {
	model := struct {
		Time qdefine.Date
	}{}
	err := json.Unmarshal([]byte(c.GetString(key)), &model)
	if err != nil {
		return 0
	}
	return model.Time
}

func (c *context) GetTime(key string) qdefine.DateTime {
	model := struct {
		Time qdefine.DateTime
	}{}
	err := json.Unmarshal([]byte(c.GetString(key)), &model)
	if err != nil {
		return 0
	}
	return model.Time
}

func (c *context) GetStruct(key string, objType reflect.Type) any {
	val := c.values.getValue(key)
	// 先转为json
	js, _ := json.Marshal(val)
	// 创建新的对象
	ptrObj := reflect.New(objType).Interface()
	// 再反转
	_ = json.Unmarshal(js, ptrObj)
	// 返回非指针对象
	obj := reflect.ValueOf(ptrObj).Elem().Interface()
	return obj
}

func (c *context) loadValues() {
	// 解析body
	contentType := c.gin.Request.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "application/json") {
		// 处理 JSON 数据
		if body, e := io.ReadAll(c.gin.Request.Body); e == nil && len(body) > 0 {
			err := c.values.loadBody(body)
			if err != nil {
				return
			}
		}
	} else if strings.HasPrefix(contentType, "multipart/form-data") {
		// 处理表单数据
		form, err := c.gin.MultipartForm()
		if err != nil {
			return
		}
		// 将非文件值加入到字典中
		for key, value := range form.Value {
			if len(value) > 0 {
				c.values.setInputValue(key, value[0])
			}
		}
	}
	// 解析Query
	querys := c.gin.Request.URL.Query()
	for k, v := range querys {
		if len(v) > 0 {
			c.values.setInputValue(k, v[0])
		}
	}
	// 解析路由参数
	for _, v := range c.gin.Params {
		c.values.setInputValue(v.Key, v.Value)
	}
	// 解析Headers
	for k, v := range c.gin.Request.Header {
		if len(v) > 0 {
			c.values.setInputValue(k, v[0])
		}
	}
}

type values struct {
	Inputs []map[string]interface{}
	//Header map[string]interface{}
}

func (d *values) loadBody(body []byte) error {
	var obj interface{}
	err := json.Unmarshal(body, &obj)
	if err != nil {
		return err
	}
	maps := make([]map[string]interface{}, 0)
	kind := reflect.TypeOf(obj).Kind()
	if kind == reflect.Slice {
		for _, o := range obj.([]interface{}) {
			maps = append(maps, o.(map[string]interface{}))
		}
	} else if kind == reflect.Map || kind == reflect.Struct {
		maps = append(maps, obj.(map[string]interface{}))
	} else {
		maps = append(maps, map[string]interface{}{"": obj})
	}
	d.Inputs = maps
	return nil
}

func (d *values) setInputValue(key string, value interface{}) {
	if len(d.Inputs) == 0 {
		d.Inputs = append(d.Inputs, map[string]interface{}{key: value})
	}
	for i := 0; i < len(d.Inputs); i++ {
		d.Inputs[i][key] = value
	}
}

func (d *values) getValue(key string) interface{} {
	if len(d.Inputs) == 0 {
		return nil
	}
	if key == "" {
		return d.Inputs[0]
	}

	var value interface{}
	if v, ok := d.Inputs[0][key]; ok {
		// 如果存在
		value = v
	} else {
		str := stringy.New(key).CamelCase()
		// 如果不存在，尝试查找
		for k, v := range d.Inputs[0] {
			if strings.ToLower(str) == strings.ToLower(stringy.New(k).CamelCase()) {
				value = v
				break
			}
		}
	}
	return value
}

func (d *values) ToMap() map[string]interface{} {
	if len(d.Inputs) == 0 {
		return nil
	}
	return d.Inputs[0]
}

func (d *values) ToMaps() []map[string]interface{} {
	return d.Inputs
}
