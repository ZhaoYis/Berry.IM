package config

import (
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// InitAppConfig 初始化配置文件
func InitAppConfig() {
	// 设置配置文件的名称
	viper.SetConfigName("app")
	// 设置配置文件的路径
	viper.AddConfigPath("config/files")
	// 设置配置文件的类型，可选，若不设置，viper 会自动检测
	viper.SetConfigType("yaml")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			// 配置文件未找到错误
			fmt.Println("Config file not found; ignore error if desired")
		}
		return
	}

	// 打印配置文件的路径
	fmt.Println("Using config file:", viper.ConfigFileUsed())

	// 监听配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}
