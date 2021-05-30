package env

import (
	"fmt"
	"os"
)

type RuntimeEnv string

// 运行环境定义
const (
	Production  RuntimeEnv = "production"
	Development RuntimeEnv = "development"
	Test        RuntimeEnv = "test"
	Unset       RuntimeEnv = ""
)

// GetEnvOrDefault 获取环境变量值。
// 成功获取到，则返回具体值；若无法获取或获取到空，则返回 defaultValue。
func GetEnvOrDefault(envName, defaultValue string) string {
	flag := ""
	v := os.Getenv(envName)
	if v == "" {
		v = defaultValue
		flag = "(default)"
	}
	fmt.Fprintf(os.Stderr, "ENV $%s: %s %s\r\n", envName, defaultValue, flag)
	return v
}

// GetRuntimeEnv 获取当前环境
// 默认为生产(production)，指定环境为开发(development)，测试(test)
func GetRuntimeEnv() RuntimeEnv {
	env := os.Getenv("CFG_ENV")
	switch env {
	case string(Production):
		return Production
	case string(Development):
		return Development
	case string(Test):
		return Test
	default:
		return Unset
	}
}