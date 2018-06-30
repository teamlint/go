package html

// Funcer 模板函数接口
type Funcer interface {
	AddFunc(funcName string, funcBody interface{})
}
