package app

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"github.com/heyu52/Fly/xlog"
	"time"
)


const HostPort = 18801
const HostPath = "/"

type App struct {
	Domain string
	Project string
	Module string
	Modules []ModuleName
	Engine *gin.Engine
	StartTime time.Time
}

var Instance *App

func (a *App) Start()  {
	port:=viper.GetInt("host.port")
	if port<1 || port>65635{
		port=HostPort
	}

	a.StartOnPort(port)
}


func (a *App) StartOnPort(port int)  {

	fmt.Printf("应用程序端口:%d\n",port)
	addr:=	fmt.Sprintf(":%v",port)

	srv:=&http.Server{
		Addr: addr,
		Handler: a.Engine,
	}

	go func() {
		err:=srv.ListenAndServe()
		if err!=nil && err!=http.ErrServerClosed{
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	// 等待中断信号以优雅地关闭服务器
	signal.Notify(quit, os.Interrupt)
	<-quit
	// 设置 5 秒的超时时间，超时后真实执行关闭
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.RegisterOnShutdown(func() {
		xlog.Info("Server exited")
	})

	//xlog.Info("Server exiting")
	if err := srv.Shutdown(ctx); err != nil {
		xlog.Fatalf("Server Shutdown: %s", err.Error())
	}
}