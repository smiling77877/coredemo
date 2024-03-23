package contract

const (
	//EnvProduction代表生产环境
	EnvProduction = "production"
	//EnvTesting代表测试环境
	EnvTesting = "testing"
	//EnvDevelopment代表开发环境
	EnvDevelopment = "development"
	//EnvKey是环境变量服务字符串凭证
	EnvKey = "hade:env"
)

// Env定义环境变量服务
type Env interface {
	//AppEnv获取当前的环境，建议分为development/testing/production
	AppEnv() string
	//IsExist判断一个环境变量是否有被设置
	IsExist(string) bool
	//Get获取某个环境变量，如果没有设置，返回""
	Get(string) string
	//All获取所有的环境变量，.env和运行环境变量融合后结果
	All() map[string]string
}
