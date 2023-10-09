package web

import (
	"encoding/base64"
	"encoding/json"
	"github.com/UritMedical/qf2/utils/launcher"
	"github.com/UritMedical/qf2/utils/qio"
	"os/exec"
)

type setting struct {
	startPort  int
	pluginPath string
}

type pluginSetting struct {
	Port       int            // 自身服务端口
	PluginPort map[string]int // 其他插件服务端口
}

func RunGinContainer() {
	launcher.Run(start, nil)
}

func start() {
	config := setting{
		startPort:  30001,
		pluginPath: "./",
	}
	// 加载插件
	dir := qio.GetFullPath(config.pluginPath)
	files, err := qio.GetFiles(dir, "plugin_*.exe")
	if err != nil {
		panic(err)
	}
	// 生成配置
	configs := loadPluginConfig(config.startPort, files)
	// 启动插件
	for file, config := range configs {
		j, _ := json.Marshal(config)
		baseConfig := base64.StdEncoding.EncodeToString(j)

		// 启动exe
		err := exec.Command(file, baseConfig).Start()
		if err != nil {
			panic(err)
		}
	}
}

func loadPluginConfig(startPort int, files []string) map[string]pluginSetting {
	configs := map[string]pluginSetting{}
	// 自动分配端口
	for _, file := range files {
		configs[file] = pluginSetting{
			Port:       startPort,
			PluginPort: map[string]int{},
		}
		startPort++
	}
	// 将其他插件的端口给每个插件
	for k, v := range configs {
		for _, file := range files {
			if file == k {
				continue
			}
			v.PluginPort[qio.GetFileNameWithoutExt((file))] = configs[file].Port
		}
	}
	return configs
}
