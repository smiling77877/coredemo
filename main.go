package main

import (
	"github.com/smiling77877/coredemo/app/console"
	"github.com/smiling77877/coredemo/app/http"
	"github.com/smiling77877/coredemo/app/provider/demo"
	"github.com/smiling77877/coredemo/framework"
	"github.com/smiling77877/coredemo/framework/provider/app"
	"github.com/smiling77877/coredemo/framework/provider/distributed"
	"github.com/smiling77877/coredemo/framework/provider/kernel"
)

func main() {
	//初始化服务容器
	container := framework.NewHadeContainer()
	//绑定App服务提供者
	container.Bind(&app.HadeAppProvider{})
	//后续初始化需要绑定的服务提供者
	container.Bind(&distributed.LocalDistributedProvider{})

	//将HTTP引擎初始化，并且作为服务提供者绑定到服务容器中
	if engine, err := http.NewHttpEngine(); err == nil {
		container.Bind(&demo.DemoProvider{})
		container.Bind(&kernel.HadeKernelProvider{HttpEngine: engine})
	}

	//运行root命令
	console.RunCommand(container)
}
