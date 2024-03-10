package framework

import (
	"log"
	"net/http"
	"strings"
)

// 框架核心架构
type Core struct {
	router      map[string]*Tree    // all routers
	middlewares []ControllerHandler // 从core这边设置的中间件
}

// 初始化框架核心架构
func NewCore() *Core {
	//初始化路由
	router := map[string]*Tree{}
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()
	return &Core{router: router}
}

// 注册中间件
func (c *Core) Use(middlewares ...ControllerHandler) {
	c.middlewares = append(c.middlewares, middlewares...)
}

// 匹配GET方法，增加路由规则
func (c *Core) Get(url string, handlers ...ControllerHandler) {
	//将core的middleware和handlers结合起来
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["GET"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

// 匹配POST方法，增加路由规则
func (c *Core) Post(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["POST"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

// 匹配PUT方法，增加路由规则
func (c *Core) Put(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["PUT"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

// 匹配DELETE方法，增加路由规则
func (c *Core) Delete(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["DELETE"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

// 从core中初始化这个Group
func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}

func (c *Core) FindRouteByRequest(r *http.Request) []ControllerHandler {
	// uri和method全部转换为大写，保证大小写不敏感
	uri := r.URL.Path
	method := r.Method
	upperMethod := strings.ToUpper(method)

	//查找第一层map
	if methodHandlers, ok := c.router[upperMethod]; ok {
		//查找第二层map
		return methodHandlers.FindHandler(uri)
	}
	return nil
}

// 所有请求都进入这个函数，这个函数负责路由分发
func (c *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//封装自定义Context
	ctx := NewContext(r, w)

	//寻找路由
	handlers := c.FindRouteByRequest(r)
	if handlers == nil {
		//如果没有找到，这里打印日志
		ctx.Json(404, "not found")
		return
	}

	//设置context中的handlers字段
	ctx.SetHandlers(handlers)

	//调用路由函数，如果返回err代表存在内部错误，返回500状态码
	if err := ctx.Next(); err != nil {
		ctx.Json(500, "inner error")
		return
	}
}
