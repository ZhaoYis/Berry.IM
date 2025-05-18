package config

import "github.com/spf13/viper"

type AppConfig struct {
	Env    string
	Server *AppServerConfig
}

type AppServerConfig struct {
	Host string
	Port string
}

// GetEnv 获取当前环境
func GetEnv() string {
	return viper.GetString("app.env")
}

func GetAppConfig() *AppConfig {
	return &AppConfig{
		Env: viper.GetString("app.env"),
		Server: &AppServerConfig{
			Host: viper.GetString("app.server.host"),
			Port: viper.GetString("app.server.port"),
		},
	}
}
