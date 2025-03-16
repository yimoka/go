# 错误处理说明

## 1. 概述
本模块提供了统一的错误处理机制，包括错误码定义、错误包装、错误转换等功能。基于 go-kratos 的错误处理系统进行了扩展，使其更适合微服务架构。

## 2. 错误类型

### 2.1 业务错误
```go
import "github.com/yimoka/go/err"

// 定义错误码
var (
    ErrUserNotFound = err.New(
        http.StatusNotFound,
        "USER_NOT_FOUND",
        "用户不存在",
    )
    
    ErrInvalidParameter = err.New(
        http.StatusBadRequest,
        "INVALID_PARAMETER",
        "无效的参数",
    )
)
```

### 2.2 系统错误
```go
var (
    ErrInternalServer = err.New(
        http.StatusInternalServerError,
        "INTERNAL_SERVER_ERROR",
        "服务器内部错误",
    )
    
    ErrServiceUnavailable = err.New(
        http.StatusServiceUnavailable,
        "SERVICE_UNAVAILABLE",
        "服务不可用",
    )
)
```

## 3. 使用方法

### 3.1 创建错误
```go
// 创建新错误
err := err.New(http.StatusBadRequest, "INVALID_EMAIL", "邮箱格式不正确")

// 包装错误
err = err.Wrap(err, "验证邮箱失败")

// 带参数的错误
err = err.NewWithData(http.StatusBadRequest, "INVALID_PARAMETER", "参数 %s 无效", "email")
```

### 3.2 错误判断
```go
// 判断错误类型
if err.Is(err, ErrUserNotFound) {
    // 处理用户不存在错误
}

// 获取错误码
code := err.Code(err)

// 获取错误信息
msg := err.Message(err)
```

### 3.3 错误转换
```go
// HTTP 错误转换
httpErr := err.FromError(err)

// gRPC 错误转换
grpcErr := err.ToGRPCError(err)
```

## 4. 错误码规范

### 4.1 错误码格式
- HTTP 状态码：标准 HTTP 状态码
- 错误码：大写字母和下划线组成
- 错误信息：清晰描述错误原因

### 4.2 常用错误码
```go
const (
    // 客户端错误 (400-499)
    BadRequest           = 400 // 请求参数错误
    Unauthorized        = 401 // 未授权
    Forbidden          = 403 // 禁止访问
    NotFound           = 404 // 资源不存在
    MethodNotAllowed   = 405 // 方法不允许
    
    // 服务端错误 (500-599)
    InternalServer     = 500 // 服务器内部错误
    NotImplemented     = 501 // 未实现
    ServiceUnavailable = 503 // 服务不可用
)
```

## 5. 最佳实践

### 5.1 错误定义
1. 错误码应该是唯一的
2. 错误信息应该清晰明确
3. 避免暴露敏感信息
4. 使用统一的错误格式

### 5.2 错误处理
1. 及时处理错误
2. 正确记录错误日志
3. 合理使用错误包装
4. 避免吞掉错误

### 5.3 错误返回
1. 统一错误响应格式
2. 合理设置 HTTP 状态码
3. 提供有用的错误信息
4. 考虑国际化需求 