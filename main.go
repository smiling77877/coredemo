package main

import (
	"github.com/smiling77877/coredemo/framework"
	"github.com/smiling77877/coredemo/framework/middleware"
	"net/http"
)

func main() {
	core := framework.NewCore()
	//core中使用use注册中间件
	//core.Use(middleware.Test1(),
	//	middleware.Test2())
	core.Use(middleware.Recovery())
	core.Use(middleware.Cost())

	//group中使用use注册中间件
	//subjectApi := core.Group("/subject")
	//subjectApi.Use(middleware.Test3())

	registerRouter(core)
	server := &http.Server{
		// 自定义的请求核心处理函数
		Handler: core,
		// 请求监听地址
		Addr: ":8080",
	}
	server.ListenAndServe()
}
