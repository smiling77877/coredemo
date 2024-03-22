package contract

import "net/http"

// KernelKey提供kernel服务凭证
const KernelKey = "hade:kernel"

// Kernel接口提供框架最核心的架构
type Kernel interface {
	//HttpEngine http.Handler结构，作为net/http框架使用，实际上是gin.Engine
	HttpEngine() http.Handler
}
