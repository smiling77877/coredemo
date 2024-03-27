package gin

import "github.com/smiling77877/coredemo/framework/contract"

// MustMakeApp从容器中获取App服务
func (c *Context) MustMakeApp() contract.App {
	return c.MustMake(contract.AppKey).(contract.App)
}

// MustMakeKernel从容器中获取Kernel服务
func (c *Context) MustMakeKernel() contract.Kernel {
	return c.MustMake(contract.KernelKey).(contract.Kernel)
}

// MustMakeConfig从容器中获取配置服务
func (c *Context) MustMakeConfig() contract.Config {
	return c.MustMake(contract.ConfigKey).(contract.Config)
}

// MustMakeLog从容器中获取日志服务
func (c *Context) MustMakeLog() contract.Log {
	return c.MustMake(contract.LogKey).(contract.Log)
}
