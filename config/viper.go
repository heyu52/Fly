package config

import "github.com/spf13/viper"

func GetStringOrDefault(key, defaultValue string) string  {
	val:=viper.GetString(key)
	if val!=""{
		return val
	}
	return  defaultValue
}
