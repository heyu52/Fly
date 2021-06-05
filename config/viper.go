package config

import "github.com/spf13/viper"

func GetStringOrDefault(key, defaultValue string) string  {
	val:=viper.GetString(key)
	if val!=""{
		return val
	}
	return  defaultValue
}

func GetBoolOrDefault(key string, defaultValue bool) bool {
	if exists := viper.IsSet(key); exists {
		return viper.GetBool(key)
	}
	return defaultValue
}

func GetIntOrDefault(key string, defaultValue int) int {
	if exists := viper.IsSet(key); exists {
		return viper.GetInt(key)
	}
	return defaultValue
}