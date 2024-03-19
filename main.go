package main

import (
	"context"
	"github.com/smiling77877/coredemo/app/provider/demo"
	"github.com/smiling77877/coredemo/framework/gin"
	"github.com/smiling77877/coredemo/framework/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//创建engine结构
	core := gin.New()
	//绑定具体的服务
	core.Bind(&demo.DemoServiceProvider{})

	core.Use(gin.Recovery())
	core.Use(middleware.Cost())

	registerRouter(core)
	server := &http.Server{
		// 自定义的请求核心处理函数
		Handler: core,
		// 请求监听地址
		Addr: ":8080",
	}
	//这个Goroutine是启动服务的Goroutine
	go func() {
		server.ListenAndServe()
	}()

	//当前的Goroutine等待信号量
	quit := make(chan os.Signal)
	//监控信号: SIGINT, SIGTERM, SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	//这里会阻塞当前Goroutine等待信号
	<-quit

	//调用Server.Shutdown来优雅结束
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
