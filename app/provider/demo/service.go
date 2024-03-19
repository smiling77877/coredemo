package demo

import (
	"github.com/smiling77877/coredemo/framework"
)

type Service struct {
	container framework.Container
}

// 初始化实例的方法
func NewService(params ...interface{}) (interface{}, error) {
	//这里需要将参数展开
	container := params[0].(framework.Container)
	//返回示例
	return &Service{container: container}, nil
}

// 实现接口
func (s *Service) GetAllstudent() []Student {
	return []Student{
		{
			ID:   1,
			Name: "foo",
		},
		{
			ID:   2,
			Name: "bar",
		},
	}
}
