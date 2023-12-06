// Package err 错误处理
package err

// Handler 定义一个错误处理器的接口
type Handler interface {
	// Handle 处理错误
	Handle(err error) error
}
