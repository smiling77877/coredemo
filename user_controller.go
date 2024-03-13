package main

import (
	"github.com/smiling77877/coredemo/framework"
	"time"
)

func UserLoginController(c *framework.Context) error {
	foo, _ := c.QueryString("foo", "def")
	//等待10s才结束运行
	time.Sleep(10 * time.Second)
	c.SetOkStatus().Json("ok, UserLoginController" + foo)
	return nil
}
