package main

import (
	"github.com/smiling77877/coredemo/framework/gin"
	"time"
)

func UserLoginController(c *gin.Context) {
	foo, _ := c.DefaultQueryString("foo", "def")
	//等待10s才结束运行
	time.Sleep(10 * time.Second)
	c.ISetOkStatus().IJson("ok, UserLoginController" + foo)
}
