package http

import (
	"github.com/smiling77877/coredemo/app/http/module/demo"
	"github.com/smiling77877/coredemo/framework/gin"
	"github.com/smiling77877/coredemo/framework/middleware/static"
)

// Routes绑定业务层路由
func Routes(r *gin.Engine) {
	// /路径先去。/dist目录下查找文件是否存在，找到使用文件服务提供服务
	r.Use(static.Serve("/", static.LocalFile("./dist", false)))
	demo.Register(r)
}
