package main

import "github.com/smiling77877/coredemo/framework"

func UserLoginController(c *framework.Context) error {
	foo, _ := c.QueryString("foo", "def")
	c.SetOkStatus().Json("ok, UserLoginController" + foo)
	return nil
}
