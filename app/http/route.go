package http

import (
	"github.com/smiling77877/coredemo/app/http/module/demo"
	"github.com/smiling77877/coredemo/framework/gin"
)

func Routes(r *gin.Engine) {
	r.Static("/dist/", "./dist/")
	demo.Register(r)
}
