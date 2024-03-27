package id

import (
	"github.com/smiling77877/coredemo/framework"
	"github.com/smiling77877/coredemo/framework/contract"
)

type HadeIDProvider struct {
}

func (provider *HadeIDProvider) Register(c framework.Container) framework.NewInstance {
	return NewHadeIDService
}

func (provider *HadeIDProvider) Boot(c framework.Container) error {
	return nil
}

func (provider *HadeIDProvider) IsDefer() bool {
	return false
}

func (provider *HadeIDProvider) Params(c framework.Container) []interface{} {
	return []interface{}{}
}

func (provider *HadeIDProvider) Name() string {
	return contract.IDKey
}
