package qconfig

import (
	"encoding/json"
	"github.com/UritMedical/qf2/utils/qio"
	"github.com/spf13/viper"
)

var (
	projectName       string   = ""              // 项目名称，如果多个服务需要共用时
	filePath          string   = "./config.yaml" // 配置文件路径
	allConfigContents [][]byte                   // 所有配置文件的内容
	hasRead           bool                       // 如果已经加载过文件，则不再加载
)

// Init
//
//	@Description: 初始化，原则上只在服务启动的时候执行一次
//	@param newProjectName 服务名称，默认为空
//	@param newFileName 配置文件路径，默认为./config.yaml
//	@param configContents 配置文件内容
func Init(newProjectName string, newFileName string, configContents [][]byte) {
	projectName = newProjectName
	if newFileName != "" {
		filePath = newFileName
	}
	allConfigContents = configContents
}

// Get
//
//	@Description: 获取配置
//	@param key 节 格式：xxx.xxx.xxx
//	@return T
func Get[T any](key string, defValue T) T {
	// 如果设置了项目名称，则追加项目名称
	if projectName != "" {
		key = projectName + "." + key
	}

	// 写入默认值
	viper.SetDefault(key, defValue)

	// 读取配置文件
	full := qio.GetFullPath(filePath)
	readConfig(full)

	// 通过json转换到t
	obj := viper.Get(key)
	js, err := json.Marshal(obj)
	if err == nil {
		newObj := new(T)
		err := json.Unmarshal(js, &newObj)
		cont := string(js)
		cont = cont + ""
		if err == nil {
			return *newObj
		}
	}

	return *new(T)
}

func readConfig(path string) {
	fullContent := make([]byte, 0)
	for _, content := range allConfigContents {
		fullContent = append(fullContent, content...)
	}
	if qio.PathExists(path) == false {
		// 文件不存在，则创建配置文件
		_ = qio.WriteAllBytes(path, fullContent, false)
		hasRead = false
	}

	// 比较内容是否有变化
	// 读取原来的文件内容
	// 替换
	// 重新生成文件

	// 读取
	if hasRead == false {
		viper.SetConfigFile(path)
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
		hasRead = true
	}
}
