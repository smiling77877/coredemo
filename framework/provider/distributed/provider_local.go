package distributed

import (
	"github.com/smiling77877/coredemo/framework"
	"github.com/smiling77877/coredemo/framework/contract"
)

// LocalDistributedProvider提供App的具体实现方法
type LocalDistributedProvider struct {
}

// Register注册HadeApp方法
func (h *LocalDistributedProvider) Register(container framework.Container) framework.NewInstance {
	return NewLocalDistributedService
}

// Boot启动调用
func (h *LocalDistributedProvider) Boot(container framework.Container) error {
	return nil
}

// IsDefer是否延迟初始化
func (h *LocalDistributedProvider) IsDefer() bool {
	return false
}

// Params获取初始化参数
func (h *LocalDistributedProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container}
}

// Name获取字符串凭证
func (h *LocalDistributedProvider) Name() string {
	return contract.DistributedKey
}
