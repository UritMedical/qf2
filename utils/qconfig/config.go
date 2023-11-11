package qconfig

import (
	"encoding/json"
	"github.com/UritMedical/qf2/utils/qio"
	"github.com/spf13/viper"
)

// Get
//
//	@Description: 获取配置
//	@param key
//	@param defValue
//	@return T
func Get[T any](key string, defValue T) T {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	// 写入默认值
	viper.SetDefault(key, defValue)

	// 获取配置
	if qio.PathExists("./config/config.yaml") {
		// 通过json转换到t
		obj := viper.Get(key)
		js, err := json.Marshal(obj)
		if err == nil {
			newObj := new(T)
			err := json.Unmarshal(js, &newObj)
			if err == nil {
				return *newObj
			}
		}
	}

	return defValue
}

// Set
//
//	@Description: 设置配置
//	@param key
//	@param value
func Set(key string, value any) {
	viper.Set(key, value)
}

// Save
//
//	@Description: 保存文件
//	@return error
func Save() error {
	return viper.WriteConfig()
}
