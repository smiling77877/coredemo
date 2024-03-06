package framework

import (
	"net/http"
	"strings"
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
	upperUrl := strings.ToUpper(url)
	c.router["GET"][upperUrl] = handler
}

func (c *Core) Post(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["POST"][upperUrl] = handler
}

func (c *Core) Put(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["PUT"][upperUrl] = handler
}

func (c *Core) Delete(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["DELETE"][upperUrl] = handler
}

func (c *Core) FindRouteByRequest(r *http.Request) ControllerHandler {
	// uri和method全部转换为大写，保证大小写不敏感
	uri := r.URL.Path
	method := r.Method
	upperMethod := strings.ToUpper(method)
	upperUri := strings.ToUpper(uri)

	//查找第一层map
	if methodHandlers, ok := c.router[upperMethod]; ok {
		//查找第二层map
		if handler, ok := methodHandlers[upperUri]; ok {
			return handler
		}
	}
	return nil
}

// 从core中初始化这个Group
func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}

// 框架核心结构实现Handler接口
func (c *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//封装自定义Context
	ctx := NewContext(r, w)
	//寻找路由
	router := c.FindRouteByRequest(r)
	if router == nil {
		//如果没有找到，这里打印日志
		ctx.Json(404, "not found")
		return
	}
	//调用路由函数，如果返回err代表存在内部错误，返回500状态码
	if err := router(ctx); err != nil {
		ctx.Json(500, "inner error")
		return
	}
}
