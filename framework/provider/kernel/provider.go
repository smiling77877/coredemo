package kernel

import (
	"github.com/smiling77877/coredemo/framework"
	"github.com/smiling77877/coredemo/framework/contract"
	"github.com/smiling77877/coredemo/framework/gin"
)

// HadeKernelProvider提供web引擎
type HadeKernelProvider struct {
	HttpEngine *gin.Engine
}

// Register注册服务提供者
func (provider *HadeKernelProvider) Register(c framework.Container) framework.NewInstance {
	return NewHadeKernelService
}

// Boot启动的时候判断是否由外界注入了Engine，如果注入的话就用注入的，如果没有就重新实例化
func (provider *HadeKernelProvider) Boot(c framework.Container) error {
	if provider.HttpEngine == nil {
		provider.HttpEngine = gin.Default()
	}
	provider.HttpEngine.SetContainer(c)
	return nil
}

// IsDefer引擎的初始化我们希望开始就进行初始化
func (provider *HadeKernelProvider) IsDefer() bool {
	return false
}

// Params参数就是一个HttpEngine
func (provider *HadeKernelProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.HttpEngine}
}

// Name提供凭证
func (provider *HadeKernelProvider) Name() string {
	return contract.KernelKey
}
