package demo

import (
	"fmt"
	"github.com/smiling77877/coredemo/framework"
)

type DemoService struct {
	// 实现接口
	Service

	//参数
	c framework.Container
}

// 初始化实例的方法
func NewDemoService(params ...interface{}) (interface{}, error) {
	//这里需要将参数展开
	c := params[0].(framework.Container)

	fmt.Println("new demo service")
	//返回示例
	return &DemoService{c: c}, nil
}

// 实现接口
func (s *DemoService) GetFoo() Foo {
	return Foo{
		Name: "i am foo",
	}
}
