package app

import (
	"errors"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/smiling77877/coredemo/framework"
	"github.com/smiling77877/coredemo/framework/util"
)

// HadeApp代表hade框架的App实现
type HadeApp struct {
	container  framework.Container //服务容器
	baseFolder string              //基础路径
	appID      string
}

func (h HadeApp) Version() string {
	return "0.0.3"
}

// BaseFolder表示基础目录，可以代表开发场景的目录，也可以代表运行时候的目录
func (h HadeApp) BaseFolder() string {
	if h.baseFolder != "" {
		return h.baseFolder
	}
	//如果参数也没有，使用默认的当前路径
	return util.GetExecDirectory()
}

func (h HadeApp) ConfigFolder() string {
	return filepath.Join(h.BaseFolder(), "config")
}

func (h HadeApp) LogFolder() string {
	return filepath.Join(h.StorageFolder(), "log")
}

func (h HadeApp) HttpFolder() string {
	return filepath.Join(h.BaseFolder(), "http")
}

func (h HadeApp) ConsoleFolder() string {
	return filepath.Join(h.BaseFolder(), "console")
}

func (h HadeApp) StorageFolder() string {
	return filepath.Join(h.BaseFolder(), "storage")
}

func (h HadeApp) ProvideFolder() string {
	return filepath.Join(h.BaseFolder(), "provider")
}

func (h HadeApp) MiddlewareFolder() string {
	return filepath.Join(h.HttpFolder(), "middleware")
}

func (h HadeApp) CommandFolder() string {
	return filepath.Join(h.ConsoleFolder(), "command")
}

func (h HadeApp) RuntimeFolder() string {
	return filepath.Join(h.StorageFolder(), "runtime")
}

func (h HadeApp) TestFolder() string {
	return filepath.Join(h.BaseFolder(), "test")
}

// NewHadeApp初始化HadeApp
func NewHadeApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}

	//有两个参数，一个是容器，一个是baseFolder
	container := params[0].(framework.Container)
	baseFolder := params[1].(string)
	appID := uuid.New().String()
	return &HadeApp{baseFolder: baseFolder, container: container, appID: appID}, nil
}
