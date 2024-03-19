package contract

// AppKey定义字符串凭证
const AppKey = "hade:app"

// App定义接口
type App interface {
	//Version定义当前版本
	Version() string
	//BaseFolder定义项目基础地址
	BaseFolder() string
	//ConfigFolder定义了配置文件路径
	ConfigFolder() string
	//LogFolder定义了日志所在路径
	LogFolder() string
	//ProvideFolder定义业务自己的服务提供者地址
	ProvideFolder() string
	//MiddlewareFolder定义业务自己的中间件
	MiddlewareFolder() string
	//CommandFolder定义业务定义的命令
	CommandFolder() string
	//RuntimeFolder定义业务的运行中间态信息
	RuntimeFolder() string
	//TestFolder存放测试所需要的信息
	TestFolder() string
}
