package launcher

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/heyu52/Fly/app"
	"github.com/heyu52/Fly/config"
	"github.com/heyu52/Fly/env"
	"github.com/heyu52/Fly/xlog"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
	"time"
)

func NewApp(cfg *config.RainbowInfo) *app.App {
	fmt.Println("应用程序准备启动.......")

	if cfg == nil {
		cfg = &config.RainbowInfo{}
	}

	err := config.Init(cfg)

	if err != nil {
		panic(errors.Wrap(err, "裙始化配置失败"))
	}

	instance := &app.App{
		Domain:    viper.GetString("app.domain"),
		Project:   viper.GetString("app.project"),
		Module:    viper.GetString("app.module"),
		StartTime: time.Now(),
	}

	app.Instance = instance

	instance.InitModule(app.ModuleLog, func() error {
		defer xlog.Sync()

		err := xlog.Init(&xlog.LogSettings{
			Level:       config.GetStringOrDefault("log.level", xlog.DefaultLevel),
			Path:        config.GetStringOrDefault("log.path", xlog.DefaultPath),
			FileName:    config.GetStringOrDefault("log.filename", xlog.DefaultFileName),
			CataLog:     config.GetStringOrDefault("log.catalog", xlog.DefaultCataLog),
			MaxFileSize: config.GetIntOrDefault("log.maxfilesize", xlog.DefaultMaxFileSize),
			MaxBackups:  config.GetIntOrDefault("log.maxbackups", xlog.DefaultMaxBackups),
			MaxAge:      config.GetIntOrDefault("log.maxage", xlog.DefaultMaxAge),
			Caller:      config.GetBoolOrDefault("log.caller", xlog.DefaultCaller),
		})
		return err
	}).
		InitModule(app.ModuleGin, func() error {
		ginMode := os.Getenv("GIN_MODE")
		if len(ginMode) == 0 {
			runtimeEnv := env.GetRuntimeEnv()
			switch runtimeEnv {
			case env.Development:
				gin.SetMode(gin.DebugMode)
			case env.Test:
				gin.SetMode(gin.TestMode)
			default:
				gin.SetMode(gin.ReleaseMode)
			}
		}
		engine := gin.New()
		engine.Use(gin.Recovery())
		xlog.Infof("%s: %s\n", gin.EnvGinMode, gin.Mode())
		app.Instance.Engine = engine
		return nil
	})

	return instance
}
