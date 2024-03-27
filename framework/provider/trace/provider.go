package trace

import (
	"github.com/smiling77877/coredemo/framework"
	"github.com/smiling77877/coredemo/framework/contract"
)

type HadeTraceProvider struct {
	c framework.Container
}

func (provider *HadeTraceProvider) Register(c framework.Container) framework.NewInstance {
	return NewHadeTraceService
}

func (provider *HadeTraceProvider) Boot(c framework.Container) error {
	provider.c = c
	return nil
}

func (provider *HadeTraceProvider) IsDefer() bool {
	return false
}

func (provider *HadeTraceProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.c}
}

func (provider *HadeTraceProvider) Name() string {
	return contract.TraceKey
}
