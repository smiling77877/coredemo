package framework

import (
	"log"
	"net/http"
)

// 框架核心架构
type Core struct {
	router map[string]map[string]ControllerHandler // 二级map
}

// 初始化框架核心架构
func NewCore() *Core {
	//定义二级map
	getRouter := map[string]ControllerHandler{}
	postRouter := map[string]ControllerHandler{}
	putRouter := map[string]ControllerHandler{}
	deleteRouter := map[string]ControllerHandler{}

	//将二级map写入一级map
	router := map[string]map[string]ControllerHandler{}
	router["GET"] = getRouter
	router["POST"] = postRouter
	router["PUT"] = putRouter
	router["DELETE"] = deleteRouter

	return &Core{router: router}
}

func (c *Core) Get(url string, handler ControllerHandler) {
}

// 框架核心结构实现Handler接口
func (c *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("core.serveHTTP")
	ctx := NewContext(r, w)

	router := c.router["foo"]
	if router == nil {
		return
	}
	log.Println("core.router")

	router(ctx)
}
