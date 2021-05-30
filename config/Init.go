package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

func Init(info *RainbowInfo) error  {
	//默认使用全局viper实例
	v := viper.GetViper()

	fmt.Printf("读取环境变量\n")
	useEnv(v)

	fmt.Printf("设置启动项的默认值\n")
	useDefault(v)

	// 启动Rainbow逻辑, 只要有一个信息填写了就被认为是要使用Rainbow的，就需要验证逻辑
	isRainbow := info != nil &&
		!(info.Url == "" && info.AppId == "" && info.Group == "" && info.UserId == "" && info.UserKey == "")

	//fmt.Printf("isRainbow:%v\n",isRainbow)

	if isRainbow == true {
		//rainbow, err := NewRainbow(info.Url, info.AppId, info.Group, info.UserId, info.UserKey)
		_, err := NewRainbow(info.Url, info.AppId, info.Group, info.UserId, info.UserKey)
		if err != nil {
			panic(errors.New("[Config Init:Rainbow New] " + err.Error()))
		}

		/*// Rainbow适配Viper框架
		err = rainbow.Adapter(v)
		if err != nil {
			panic(errors.New("[Config Init:Rainbow Adapter] " + err.Error()))
		}

		//默认开启回调监控
		err = rainbow.Watch()
		if err != nil {
			panic(errors.New("[Config Init:Rainbow Watch] " + err.Error()))
		}*/
	} else {

		fmt.Printf("读取配置文件:config.yaml\n")
		// 如果Rainbow没有成功开启，扫描和读取本地文件
		useConfigFile(v)
	}

	return nil
}