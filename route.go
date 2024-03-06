package main

import "github.com/smiling77877/coredemo/framework"

func registerRouter(core *framework.Core) {
	//需求1，2:HTTP方法+静态路由匹配
	core.Get("/user/login", UserLoginController)

	//需求3:批量通用前缀
	subjectApi := core.Group("/subject")
	{
		//需求4：动态路由
		subjectApi.Get("/list", SubjectListController)
	}
}
