package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/heyu52/Fly/config"
	"github.com/heyu52/Fly/xlog"
	"github.com/pkg/errors"
	"time"
)

type ModuleName string

const (
	ModuleConfig   ModuleName = "config"
	ModuleLog      ModuleName = "log"
	ModuleMonitor  ModuleName = "monitor"
	ModuleGin      ModuleName = "gin"
	ModuleSysAPI   ModuleName = "sys-api"
	ModuleRedis    ModuleName = "redis"
	ModuleOrm      ModuleName = "orm"
	ModuleATQ      ModuleName = "atq"
	ModuleMyOA     ModuleName = "myoa"
	ModuleWorkflow ModuleName = "workflow"
	ModuleTDMQ     ModuleName = "tdmq"
	ModuleI18N     ModuleName = "i18n"

	MWTracer    ModuleName = "mw.tracer"
	MWXHTTP     ModuleName = "mw.xhttp"
	MWIp        ModuleName = "mw.ip"
	MWSignature ModuleName = "mw.signature"
	MWAuth      ModuleName = "mw.auth"
	MWQpsLimit  ModuleName = "mw.qps-limit"
)

func (a *App) PubApi(generatedFn func(eng *gin.Engine, hostPath string) error) *App {
	a.InitModule(ModuleName("pubapi"), func() error {
		hostPath := config.GetStringOrDefault("host.path", HostPath)
		return generatedFn(a.Engine, hostPath)
	})
	return a
}

func (a *App) InitModule(moduleName ModuleName, initFn func() error) *App {
	var err error
	t1 := time.Now()
	defer func() {
		duration := time.Now().Sub(t1)

		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = fmt.Errorf("unknown error from InitModule: %v", x)
			}
		}

		if err != nil {
			panic(errors.Wrapf(err, "[%s] initialize failed (%dms)", moduleName, duration))
		}

		xlog.Infof("[%s] initialized (%dms)", moduleName, duration.Nanoseconds()/1000000)
		a.Modules = append(a.Modules, moduleName)
	}()

	xlog.Infof("[%s] initialing...", moduleName)
	err = initFn()

	return a
}
