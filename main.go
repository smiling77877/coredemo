package main

import (
	"context"
	"github.com/smiling77877/coredemo/framework"
	"github.com/smiling77877/coredemo/framework/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	core := framework.NewCore()
	core.Use(middleware.Recovery())
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
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
