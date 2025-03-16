# yimoko-go

## 1. 项目介绍
基于 go-kratos 框架开发微服务的脚手架，集成了常用的中间件，方便快速开发微服务。本项目提供了一套完整的微服务开发框架，包含了项目开发中常用的功能组件和最佳实践。

## 2. 技术栈
- 框架：[go-kratos](https://go-kratos.dev/) v2.8.2
- 数据库：支持 MySQL、PostgreSQL
- 缓存：Redis
- 链路追踪：OpenTelemetry
- ORM：[ent](https://entgo.io/)
- 工作流引擎：Temporal
- 国际化：go-i18n
- 其他：protobuf、grpc

## 3. 主要特性
- 完整的微服务架构支持
- HTTP/gRPC 服务支持
- 分布式链路追踪
- 统一的错误处理
- 多语言支持
- 缓存中间件
- 数据库 ORM
- 配置中心
- 日志管理
- 工作流支持

## 4. 目录结构
```
.
├── app/            # 应用核心代码
├── cache/          # 缓存相关代码
├── config/         # 配置文件及生成的 proto
├── data/          # 数据层代码
├── ent/           # ent ORM 相关代码
├── err/           # 错误处理
├── lang/          # 国际化资源
├── logger/        # 日志处理
├── middleware/    # 中间件
│   ├── meta/     # 元数据中间件
│   └── trace/    # 链路追踪中间件
├── server/        # 服务器相关代码
├── third_party/   # 第三方依赖
├── trace/         # 链路追踪
├── utils/         # 工具函数
└── workflow/      # 工作流相关代码
```

## 5. 快速开始

### 5.1 环境要求
- Go 1.22+
- MySQL 5.7+ 或 PostgreSQL
- Redis 6.0+

### 5.2 安装
```bash
go get github.com/yimoka/go
```

## 6. 使用说明
详细的使用说明和示例请参考各模块的文档：
- [配置说明](./config/README.md)
- [中间件使用](./middleware/README.md)
- [错误处理](./err/README.md)
- [国际化使用](./lang/README.md)

## 7. 贡献指南
欢迎提交 Issue 或 Pull Request。

## 8. 开源协议
本项目采用 MIT 协议开源。