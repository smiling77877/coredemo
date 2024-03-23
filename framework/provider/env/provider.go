package env

import (
	"github.com/smiling77877/coredemo/framework"
	"github.com/smiling77877/coredemo/framework/contract"
)

type HadeEnvProvider struct {
	Folder string
}

func (provider *HadeEnvProvider) Register(c framework.Container) framework.NewInstance {
	return NewHadeEnv
}

func (provider *HadeEnvProvider) Boot(c framework.Container) error {
	app := c.MustMake(contract.AppKey).(contract.App)
	provider.Folder = app.BaseFolder()
	return nil
}

func (provider *HadeEnvProvider) IsDefer() bool {
	return false
}

func (provider *HadeEnvProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.Folder}
}

func (provider *HadeEnvProvider) Name() string {
	return contract.EnvKey
}
