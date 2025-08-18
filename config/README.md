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

## 指标配置 (Metrics)

指标配置用于 OpenTelemetry 指标收集和上报。

### 配置字段说明

- `endpoint`: 指标收集器的地址和端口，格式为 `host:port`
  - 如果未配置，将使用默认值 `apmplus-cn-beijing.ivolces.com:4317`
  - 系统会记录警告日志提示使用默认值
- `service`: 服务名称，用于标识不同的服务
- `env`: 环境标识，如 `development`、`staging`、`production`
- `namespace`: 命名空间，用于组织指标
- `subsystem`: 子系统，用于进一步分类指标
- `insecure`: 是否使用不安全连接（TLS），生产环境建议设置为 `false`
- `labels`: 自定义标签，会添加到所有指标上
- `headers`: 自定义请求头，用于认证或其他目的
  - 如果未配置，将使用默认的 `X-ByteAPM-AppKey`
  - 系统会记录日志显示使用的请求头数量

### 配置优先级

1. **配置文件中的值** - 最高优先级
2. **默认值** - 当配置缺失时使用
3. **系统自动检测** - 自动选择最佳配置

### 配置示例

**最小配置（使用默认值）：**
```yaml
metrics:
  service: "my-service"
  env: "production"
```

**完整配置：**
```yaml
metrics:
  endpoint: "your-metrics-endpoint:4317"
  service: "my-service"
  env: "production"
  namespace: "my-namespace"
  subsystem: "api"
  insecure: true
  labels:
    version: "v1.0.0"
    region: "cn-beijing"
  headers:
    X-ByteAPM-AppKey: "your-app-key"
    Authorization: "Bearer your-token"
```

### 故障排除

如果遇到连接错误，请检查：

1. **网络连接**: 确保应用服务器能够访问指标收集器
2. **防火墙设置**: 检查防火墙是否阻止了 4317 端口
3. **认证信息**: 验证 headers 中的认证信息是否正确
4. **超时设置**: 如果网络延迟较高，可以调整超时时间
5. **重试机制**: 代码中已内置重试机制，会自动处理临时连接问题

### 默认值

如果没有配置某些字段，系统会使用以下默认值：

- `endpoint`: `apmplus-cn-beijing.ivolces.com:4317`
- `headers`: 包含默认的 `X-ByteAPM-AppKey`
- `timeout`: 30 秒
- `retry`: 启用重试，初始间隔 1 秒，最大间隔 5 秒，最大重试时间 30 秒 

### 如何判断指标上传成功

#### 1. 日志监控
系统会自动记录指标上传状态：
- 成功：`指标上传成功 [endpoint]`
- 失败：`指标上传失败 [endpoint]: error_message`

#### 2. 程序化检查
```go
import "github.com/yimoka/go/metrics"

// 检查指标上传状态
if err := metrics.CheckMetricsUploadStatus(); err != nil {
    log.Printf("指标上传失败: %v", err)
} else {
    log.Printf("指标上传成功")
}

// 获取详细状态信息
status := metrics.GetDetailedMetricsStatus()
log.Printf("状态: %+v", status)

// 健康检查
health := metrics.HealthCheck()
log.Printf("健康状态: %+v", health)
```

#### 3. HTTP 健康检查端点
可以创建健康检查端点来监控指标上传状态：
```go
func healthHandler(w http.ResponseWriter, r *http.Request) {
    health := metrics.HealthCheck()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(health)
}
```

#### 4. 监控指标
系统会定期（每30秒）自动检查指标上传状态，并在日志中记录结果。

#### 5. 手动刷新
```go
// 手动刷新指标（用于测试或调试）
if err := metrics.ManualFlush(); err != nil {
    log.Printf("手动刷新失败: %v", err)
} else {
    log.Printf("手动刷新成功")
}
```

#### 6. 指标内容打印
系统会自动打印详细的指标上传内容：

**上传成功时显示：**
```
✅ 指标上传成功 [endpoint] - 2025-01-19 10:30:00
   上传内容:
   - 服务指标: HTTP请求计数、响应时间等
   - 业务指标: 自定义业务指标
   - 系统指标: 资源使用情况
   - 标签信息: 服务名称、环境、版本等
```

**上传失败时显示：**
```
❌ 指标上传失败 [endpoint] - 2025-01-19 10:30:00
   错误详情: connection refused
   建议检查:
   - 网络连接是否正常
   - 认证信息是否正确
   - 防火墙设置
   - 指标收集器是否运行
```

**详细指标内容：**
```go
// 打印当前指标内容详情
metrics.PrintMetricsContent()

// 输出示例：
=== 当前指标内容详情 ===
时间: 2025-01-19 10:30:00
服务: your-service-name
配置信息:
  - 指标收集器: OTLP gRPC
  - 收集间隔: 10秒
  - 超时设置: 3秒
  - 重试机制: 已启用
资源信息:
  - 服务名称: your-service-name
  - 环境标识: 已配置
  - 自定义标签: 已配置
连接信息:
  - 连接状态: 活跃
  - 最后检查: 10:30:00
  - 下次检查: 10:30:30
========================
``` 