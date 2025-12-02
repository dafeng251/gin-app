package initialize

import (
	"fmt"
	"gin-app/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitConfig() {
	v := viper.New()
	v.SetConfigFile("config.yaml") // 指定配置文件路径
	v.SetConfigType("yaml")        // 指定文件类型

	// 1. 读取配置
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 2. 将读取的配置映射到全局变量 global.CONF 中
	if err := v.Unmarshal(&global.CONF); err != nil {
		panic(fmt.Errorf("Unmarshal config failed: %s \n", err))
	}

	// 3. 监听配置文件变化 (热加载)
	// 当你修改 config.yaml 并保存时，程序会自动更新配置，无需重启
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件被修改:", e.Name)
		if err := v.Unmarshal(&global.CONF); err != nil {
			fmt.Println("配置重载失败:", err)
		}
	})
}
