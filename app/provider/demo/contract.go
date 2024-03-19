package demo

// Demo服务的key
const DemoKey = "demo"

// Demo服务的接口
type IService interface {
	GetAllstudent() []Student
}

// Demo服务接口定义的一个数据结构
type Student struct {
	ID   int
	Name string
}
