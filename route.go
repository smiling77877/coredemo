package main

import (
	"github.com/smiling77877/coredemo/framework/gin"
	"github.com/smiling77877/coredemo/framework/middleware"
)

func registerRouter(core *gin.Engine) {
	//需求1，2:HTTP方法+静态路由匹配
	core.GET("/user/login", middleware.Test3(), UserLoginController)

	//需求3:批量通用前缀
	subjectApi := core.Group("/subject")
	{
		subjectApi.Use(middleware.Test3())
		//需求4：动态路由
		subjectApi.DELETE("/:id", SubjectDelController)
		subjectApi.PUT("/:id", SubjectUpdateController)
		subjectApi.GET("/:id", middleware.Test3(), SubjectGetController)
		subjectApi.GET("/list/all", SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.GET("/name", SubjectNameController)
		}
	}
}
