# 配置说明

## 1. 概述
本模块基于 go-kratos 的配置系统，提供了统一的配置管理功能。配置文件使用 Protocol Buffers 定义，支持多种配置源。

## 2. 配置结构
配置文件使用 protobuf 定义，主要包含以下部分：
- 服务基础配置
- 数据库配置
- Redis 配置
- 链路追踪配置
- 日志配置
- 中间件配置

## 3. 使用方法

### 3.1 配置定义
在 `config.proto` 中定义配置结构：
```protobuf
message Bootstrap {
    Server server = 1;
    Data data = 2;
    Trace trace = 3;
    Logger logger = 4;
}
```

### 3.2 加载配置
```go
import "github.com/yimoka/go/config"

// 加载配置
conf, err := config.Load("config.yaml")
if err != nil {
    panic(err)
}
```

### 3.3 配置热更新
支持配置热更新，当配置文件发生变化时，会自动重新加载：
```go
// 监听配置变化
config.Watch(func(conf *conf.Bootstrap) {
    // 处理配置变更
})
```

## 4. 配置项说明

### 4.1 服务配置
```yaml
server:
  http:
    addr: 0.0.0.0:8000
    timeout: 1s
  grpc:
    addr: 0.0.0.0:9000
    timeout: 1s
```

### 4.2 数据库配置
```yaml
data:
  database:
    driver: mysql
    source: root:password@tcp(127.0.0.1:3306)/test
  redis:
    addr: 127.0.0.1:6379
    password: ""
    db: 0
```

### 4.3 链路追踪配置
```yaml
trace:
  endpoint: http://localhost:14268/api/traces
  sampler: 1.0
  batcher: jaeger
```

### 4.4 日志配置
```yaml
logger:
  level: info
  encoding: json
  output: stdout
```

## 5. 最佳实践
1. 不同环境使用不同的配置文件
2. 敏感信息使用环境变量注入
3. 合理使用配置默认值
4. 注意配置的版本管理 