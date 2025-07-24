# 日志过滤功能

本模块基于 [Kratos 框架](https://go-kratos.dev/docs/component/log#filter-%E6%97%A5%E5%BF%97%E8%BF%87%E6%BB%A4) 实现了完整的日志过滤功能，支持按级别过滤、按字段脱敏和按值脱敏。

## 功能特性

### 1. 按日志级别过滤
可以设置最低日志级别，低于该级别的日志将被过滤掉。

### 2. 按字段名脱敏
可以指定敏感字段名，这些字段的值将被 `***` 替换。

### 3. 按值脱敏
可以指定敏感值，匹配的值将被 `***` 替换。

### 4. 组合过滤
支持同时使用多种过滤方式。

## 配置示例

### 配置文件 (config.yaml)

```yaml
logger:
  provider: "std"  # 日志提供商: std, tencent, otel
  filterLevel: "warn"  # 只输出 warn 及以上级别的日志
  filterKeys: 
    - "password"
    - "token"
    - "secret"
    - "api_key"
  filterValues:
    - "admin123"
    - "sensitive_data"
    - "confidential"
  alsoStd: true  # 是否同时输出到标准输出
```

### 配置文件 (config.proto)

```protobuf
message Logger {
  string provider = 1;
  string filterLevel = 12;  // 日志过滤的级别: debug, info, warn, error, fatal
  repeated string filterKeys = 10;  // 日志过滤的 keys (脱敏字段名)
  repeated string filterValues = 11;  // 日志过滤的 values (脱敏值)
  bool alsoStd = 9;  // 是否同时输出到标准输出
}
```

## 使用示例

### 1. 基本使用

```go
package main

import (
    "github.com/go-kratos/kratos/v2/log"
    "github.com/yimoka/go/config"
    "github.com/yimoka/go/logger"
)

func main() {
    // 创建配置
    conf := &config.Config{
        Server: &config.Server{
            Id:      "my-service",
            Name:    "user-service",
            Version: "1.0.0",
        },
        Logger: &config.Logger{
            FilterLevel:  "warn",
            FilterKeys:   []string{"password", "token"},
            FilterValues: []string{"admin123"},
        },
    }

    // 获取logger
    logger := logger.GetLogger(conf)
    helper := log.NewHelper(logger)

    // 这些日志会被过滤掉（级别低于warn）
    helper.Debug("debug message")
    helper.Info("info message")

    // 这些日志会正常输出
    helper.Warn("warning message")
    helper.Error("error message")

    // 敏感字段会被脱敏
    helper.Infow("user login", "password", "mypassword123")  // password=***
    helper.Infow("api call", "token", "jwt_token_here")      // token=***
}
```

### 2. 不同日志提供商

```go
// 标准输出
conf.Logger.Provider = "std"

// 腾讯云日志
conf.Logger.Provider = "tencent"
conf.Logger.TopicID = "your-topic-id"
conf.Logger.AccessKey = "your-access-key"
conf.Logger.AccessSecret = "your-access-secret"
conf.Logger.Endpoint = "ap-guangzhou.cls.tencentcloudapi.com"

// OpenTelemetry
conf.Logger.Provider = "otel"
conf.Logger.Endpoint = "localhost:4317"
conf.Logger.Token = "your-token"  // 可选
```

### 3. 高级过滤配置

```go
conf := &config.Config{
    Logger: &config.Logger{
        // 只输出错误级别日志
        FilterLevel: "error",
        
        // 脱敏敏感字段
        FilterKeys: []string{
            "password",
            "token", 
            "secret",
            "api_key",
            "private_key",
            "credit_card",
            "phone",
            "email",
        },
        
        // 脱敏敏感值
        FilterValues: []string{
            "admin123",
            "password123",
            "secret_key",
            "confidential",
            "internal_only",
        },
        
        // 同时输出到标准输出
        AlsoStd: true,
    },
}
```

## 日志级别说明

- `debug`: 调试信息
- `info`: 一般信息
- `warn`: 警告信息
- `error`: 错误信息
- `fatal`: 致命错误（会中断程序）

## 脱敏效果示例

### 输入日志
```go
helper.Infow("user login", 
    "user_id", "12345",
    "password", "mypassword123",
    "token", "jwt_token_here",
    "action", "login",
    "code", "admin123",
)
```

### 输出日志（应用过滤后）
```
INFO ts=2025-07-24T04:28:27+08:00 caller=main.go:25 service.id=my-service service.name=user-service service.version=1.0.0 trace_id= span_id= msg=user login user_id=12345 password=*** token=*** action=login code=***
```

## 注意事项

1. **过滤级别**: 设置 `filterLevel` 后，低于该级别的日志将被完全过滤，不会输出到任何地方。

2. **脱敏字段**: `filterKeys` 中的字段名会进行精确匹配，匹配的字段值会被替换为 `***`。

3. **脱敏值**: `filterValues` 中的值会进行精确匹配，匹配的值会被替换为 `***`。

4. **组合使用**: 可以同时使用多种过滤方式，它们会按顺序应用。

5. **性能影响**: 过滤功能对性能影响很小，但建议合理配置过滤规则。

## 测试

运行测试来验证过滤功能：

```bash
go test ./logger -v
```

测试包括：
- 按级别过滤测试
- 按字段脱敏测试
- 按值脱敏测试
- 组合过滤测试
- 敏感数据脱敏测试 