package contract

import "time"

// DistributedKey定义字符串凭证
const DistributedKey = "hade:distributed"

// Distributed分布式服务
type Distributed interface {
	//Select分布式选择器，所有节点对某个服务进行抢占，只选择其中一个节点
	//ServiceName服务名字
	//appID当前的AppID
	//holdTime分布式选择器hold住的时间
	//返回值
	//selectAppID分布式选择器最终选择的App
	//err异常才返回，如果没有被选择，不返回err
	Select(serviceName string, appID string, holdTime time.Duration) (selectAppID string, err error)
}
