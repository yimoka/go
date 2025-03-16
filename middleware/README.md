# 中间件使用说明

## 1. 概述
本项目提供了一系列常用的中间件，用于处理跨横切关注点（Cross-cutting Concerns）的功能，如链路追踪、日志记录、错误处理等。

## 2. 可用中间件

### 2.1 元数据中间件 (meta)
用于处理请求上下文中的元数据信息。

```go
import "github.com/yimoka/go/middleware/meta"

// 使用元数据中间件
srv.Use(meta.Server())
```

主要功能：
- 传递请求头信息
- 处理用户认证信息
- 处理请求追踪 ID
- 处理语言设置

### 2.2 链路追踪中间件 (trace)
集成 OpenTelemetry 的分布式链路追踪功能。

```go
import "github.com/yimoka/go/middleware/trace"

// 使用链路追踪中间件
srv.Use(trace.Server())
```

主要功能：
- 自动生成 Trace ID
- 记录请求耗时
- 记录请求参数
- 记录错误信息

## 3. 使用方法

### 3.1 HTTP 服务中使用
```go
import (
    "github.com/yimoka/go/server"
    "github.com/yimoka/go/middleware/meta"
    "github.com/yimoka/go/middleware/trace"
)

func NewHTTPServer(c *conf.Server) *http.Server {
    srv := server.NewHTTPServer(c)
    
    // 注册中间件
    srv.Use(
        meta.Server(),
        trace.Server(),
    )
    
    return srv
}
```

### 3.2 gRPC 服务中使用
```go
import (
    "github.com/yimoka/go/server"
    "github.com/yimoka/go/middleware/meta"
    "github.com/yimoka/go/middleware/trace"
)

func NewGRPCServer(c *conf.Server) *grpc.Server {
    srv := server.NewGRPCServer(c)
    
    // 注册中间件
    srv.Use(
        meta.Server(),
        trace.Server(),
    )
    
    return srv
}
```

## 4. 自定义中间件
您可以通过实现 middleware.Middleware 接口来创建自定义中间件：

```go
type Middleware interface {
    Handler(handler middleware.Handler) middleware.Handler
}
```

示例：
```go
func CustomMiddleware() middleware.Middleware {
    return func(handler middleware.Handler) middleware.Handler {
        return func(ctx context.Context, req interface{}) (interface{}, error) {
            // 前置处理
            
            resp, err := handler(ctx, req)
            
            // 后置处理
            
            return resp, err
        }
    }
}
```

## 5. 最佳实践
1. 合理安排中间件顺序
2. 避免中间件中进行耗时操作
3. 正确处理错误传递
4. 注意中间件的性能影响 