package main

import (
	"fmt"
	"github.com/smiling77877/coredemo/app/provider/demo"
	"github.com/smiling77877/coredemo/framework/gin"
)

func SubjectAddController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectAddController")
}

// 对应路由 /subject/list/all
func SubjectListController(c *gin.Context) {
	//获取demo服务实例
	demoService := c.MustMake(demo.Key).(demo.Service)

	//调用服务实例的方法
	foo := demoService.GetFoo()

	c.ISetOkStatus().IJson(foo)
}

func SubjectDelController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectDelController")
}

func SubjectUpdateController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectUpdateController")
}

func SubjectGetController(c *gin.Context) {
	subjectId, _ := c.DefaultParamInt("id", 0)
	c.ISetOkStatus().IJson("ok, SubjectGetController" + fmt.Sprint(subjectId))
}

func SubjectNameController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectNameController")
}
