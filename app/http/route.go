package http

import (
	"github.com/smiling77877/coredemo/app/http/module/demo"
	"github.com/smiling77877/coredemo/framework/gin"
	"github.com/smiling77877/coredemo/framework/middleware"
)

// Routes绑定业务层路由
func Routes(r *gin.Engine) {
	r.Static("/dist/", "./dist/")

	r.Use(middleware.Trace())
	demo.Register(r)
}
