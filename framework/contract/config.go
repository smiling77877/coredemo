package contract

import "time"

// Config定义了配置文件服务，读取配置文件，支持点分割的路径读取
// 例如：.Get("app.name")表示从app文件中读取name属性
// 建议使用yaml属性
type Config interface {
	//IsExist检查一个属性是否存在
	IsExist(key string) bool
	//Get获取一个属性值
	Get(key string) interface{}
	//GetBool获取一个bool属性
	GetBool(key string) bool
	//GetInt获取一个int属性
	GetInt(key string) int
	//GetFloat64获取一个float64属性
	GetFloat64(key string) float64
	//GetTime获取一个time属性
	GetTime(key string) time.Time
	//GetString获取一个string属性
	GetString(key string) string
	//GetIntSlice获取一个int切片属性
	GetIntSlice(key string) []int
	//GetStringSlice获取一个string切片
	GetStringSlice(key string) []string
	//GetStringMap获取一个string为key,interface为val的map
	GetStringMap(key string) map[string]interface{}
	//GetStringMapString获取一个string为key,string为val的map
	GetStringMapString(key string) map[string]string
	//GetStringMapStringSlice获取一个string为key,string切片为val的map
	GetStringMapStringSlice(key string) map[string][]string
	//Load加载配置到某个对象
	Load(key string, val interface{}) error
}
