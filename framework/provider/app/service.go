package app

import (
	"errors"
	"flag"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/smiling77877/coredemo/framework"
	"github.com/smiling77877/coredemo/framework/util"
)

// HadeApp代表hade框架的App实现
type HadeApp struct {
	container  framework.Container //服务容器
	baseFolder string              //基础路径
	appID      string              //表示当前这个app的唯一id,可以用于分布式锁等
	configMap  map[string]string   //配置加载
}

// AppID表示这个App的唯一ID
func (app HadeApp) AppID() string {
	return app.appID
}

// Version实现版本
func (app HadeApp) Version() string {
	return "0.0.3"
}

// BaseFolder表示基础目录，可以代表开发场景的目录，也可以代表运行时候的目录
func (app HadeApp) BaseFolder() string {
	if app.baseFolder != "" {
		return app.baseFolder
	}
	//如果参数也没有，使用默认的当前路径
	return util.GetExecDirectory()
}

// ConfigFolder表示配置文件地址
func (app HadeApp) ConfigFolder() string {
	if val, ok := app.configMap["config_folder"]; ok {
		return val
	}
	return filepath.Join(app.BaseFolder(), "config")
}

// LogFolder表示日志存放地址
func (app HadeApp) LogFolder() string {
	if val, ok := app.configMap["log_folder"]; ok {
		return val
	}
	return filepath.Join(app.StorageFolder(), "log")
}

func (app HadeApp) HttpFolder() string {
	if val, ok := app.configMap["http_folder"]; ok {
		return val
	}
	return filepath.Join(app.BaseFolder(), "app", "http")
}

func (app HadeApp) ConsoleFolder() string {
	if val, ok := app.configMap["console_folder"]; ok {
		return val
	}
	return filepath.Join(app.BaseFolder(), "app", "console")
}

func (app HadeApp) StorageFolder() string {
	if val, ok := app.configMap["storage_folder"]; ok {
		return val
	}
	return filepath.Join(app.BaseFolder(), "storage")
}

// ProvideFolder定义业务自己的服务提供者地址
func (app HadeApp) ProvideFolder() string {
	if val, ok := app.configMap["provider_folder"]; ok {
		return val
	}
	return filepath.Join(app.BaseFolder(), "app", "provider")
}

// MiddlewareFolder定义业务自己定义的中间件
func (app HadeApp) MiddlewareFolder() string {
	if val, ok := app.configMap["middleware_folder"]; ok {
		return val
	}
	return filepath.Join(app.HttpFolder(), "middleware")
}

// CommandFolder定义业务定义的命令
func (app HadeApp) CommandFolder() string {
	if val, ok := app.configMap["command_folder"]; ok {
		return val
	}
	return filepath.Join(app.ConsoleFolder(), "command")
}

// RuntimeFolder定义业务的运行中间态信息
func (app HadeApp) RuntimeFolder() string {
	if val, ok := app.configMap["runtime_folder"]; ok {
		return val
	}
	return filepath.Join(app.StorageFolder(), "runtime")
}

// TestFolder定义测试需要的信息
func (app HadeApp) TestFolder() string {
	if val, ok := app.configMap["test_folder"]; ok {
		return val
	}
	return filepath.Join(app.BaseFolder(), "test")
}

// NewHadeApp初始化HadeApp
func NewHadeApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}

	//有两个参数，一个是容器，一个是baseFolder
	container := params[0].(framework.Container)
	baseFolder := params[1].(string)
	//如果没有设置，则使用参数
	if baseFolder == "" {
		flag.StringVar(&baseFolder, "base_folder", "", "base_folder参数，默认为当前路径")
		flag.Parse()
	}
	appID := uuid.New().String()
	configMap := map[string]string{}
	return &HadeApp{baseFolder: baseFolder, container: container, appID: appID, configMap: configMap}, nil
}

// LoadAppConfig加载配置map
func (app *HadeApp) LoadAppConfig(kv map[string]string) {
	for key, val := range kv {
		app.configMap[key] = val
	}
}
