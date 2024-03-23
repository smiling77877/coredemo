package cobra

import "github.com/smiling77877/coredemo/framework/contract"

// MustMakeApp 从容器中获取App服务
func (c *Command) MustMakeApp() contract.App {
	return c.GetContainer().MustMake(contract.AppKey).(contract.App)
}

// MustMakeKernel从容器中获取Kernel服务
func (c *Command) MustMakeKernel() contract.Kernel {
	return c.GetContainer().MustMake(contract.KernelKey).(contract.Kernel)
}
