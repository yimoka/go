# 国际化使用说明

## 1. 概述
本模块基于 go-i18n 实现了国际化支持，提供了多语言翻译、语言切换、消息格式化等功能。支持多种语言资源文件格式，便于管理和维护。

## 2. 功能特性
- 支持多语言切换
- 支持消息模板
- 支持复数形式
- 支持语言回退
- 支持动态加载
- 支持 YAML/JSON 格式

## 3. 使用方法

### 3.1 初始化
```go
import "github.com/yimoka/go/lang"

// 初始化国际化
i18n := lang.New(lang.Config{
    DefaultLanguage: "zh-CN",
    FallbackLanguage: "en-US",
    ResourcesPath: "resources/lang",
})
```

### 3.2 定义翻译
在 `resources/lang` 目录下创建语言文件：

zh-CN.yaml:
```yaml
welcome:
  message: "欢迎使用我们的服务"
user:
  created: "用户 {{.Name}} 创建成功"
  deleted: "用户已删除"
errors:
  not_found: "找不到资源"
  internal: "服务器内部错误"
```

en-US.yaml:
```yaml
welcome:
  message: "Welcome to our service"
user:
  created: "User {{.Name}} created successfully"
  deleted: "User deleted"
errors:
  not_found: "Resource not found"
  internal: "Internal server error"
```

### 3.3 使用翻译
```go
// 简单翻译
msg := i18n.Translate("welcome.message")

// 带参数翻译
msg = i18n.Translate("user.created", map[string]interface{}{
    "Name": "John",
})

// 切换语言
i18n.SetLanguage("en-US")
msg = i18n.Translate("welcome.message")
```

### 3.4 HTTP 中间件
```go
import "github.com/yimoka/go/middleware/i18n"

// 使用国际化中间件
srv.Use(i18n.Server())
```

### 3.5 gRPC 拦截器
```go
import "github.com/yimoka/go/middleware/i18n"

// 使用国际化拦截器
srv.Use(i18n.UnaryServerInterceptor())
```

## 4. 高级功能

### 4.1 复数形式
```yaml
items:
  count:
    one: "{{.Count}} 个项目"
    other: "{{.Count}} 个项目"
```

```go
msg := i18n.TranslatePlural("items.count", 2, map[string]interface{}{
    "Count": 2,
})
```

### 4.2 消息模板
```yaml
greeting:
  morning: "早上好，{{.Name}}"
  afternoon: "下午好，{{.Name}}"
  evening: "晚上好，{{.Name}}"
```

```go
msg := i18n.Translate("greeting.morning", map[string]interface{}{
    "Name": "张三",
})
```

### 4.3 语言检测
```go
// 从请求头检测语言
lang := i18n.DetectLanguage(r.Header.Get("Accept-Language"))

// 设置语言
i18n.SetLanguage(lang)
```

## 5. 最佳实践

### 5.1 翻译文件组织
1. 按模块分类
2. 使用层级结构
3. 保持键名一致
4. 注意文案规范

### 5.2 开发建议
1. 使用常量定义键名
2. 及时更新所有语言版本
3. 注意特殊字符转义
4. 考虑文本长度差异

### 5.3 性能优化
1. 缓存翻译结果
2. 按需加载语言包
3. 合理使用模板
4. 避免频繁切换语言 