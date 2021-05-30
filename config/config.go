package config

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type RainbowInfo struct {
	Url string
	AppId string
	UserId string
	Group string
	UserKey string
}

// useEnv 允许通过环境变量进行设置
func useEnv(v *viper.Viper) {
	// 通过环境变量设置时，“host.path” 应当设置为 “HOST_PATH”
	//envReplacer := strings.NewReplacer(".", "_")
	//v.SetEnvKeyReplacer(envReplacer)
	v.AutomaticEnv()
}

// 默认配置
func useDefault(v *viper.Viper) {
	v.SetDefault("host.path", "/")
	v.SetDefault("host.port", 80)
	v.SetDefault("log.level", "info")
	v.SetDefault("log.path", "./logs")
	v.SetDefault("log.filename", "app.log")
	v.SetDefault("log.maxfilesize", 1024)
	v.SetDefault("log.maxbackups", 3)
	v.SetDefault("log.maxage", 7)
	v.SetDefault("xhttp.log2xx", true)
	v.SetDefault("serve.log2xx", true)
	v.SetDefault("serve.dump.request.cookie", false)
	v.SetDefault("serve.dump.response.cookie", false)
}

// 支持本地文件config.yaml，但不支持本地文件变更监听
func useConfigFile(v *viper.Viper) {
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config/")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		fmt.Errorf("[Config Init: Configfile] Fatal error config file: %s \n", err)
	}
}

// ReadKey 根据节点名称从配置读取结构体
func ReadKey(key string, obj interface{}) {
	err := viper.UnmarshalKey(key, &obj)
	if err != nil {
		panic(errors.Wrapf(err, "unable to decode '%s' into struct", key))
	}
}