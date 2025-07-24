// Package app kratos.go
package app

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	conf "github.com/yimoka/go/config"
)

// DepsInit 依赖初始化 当前仅用于占位 用于当依赖注入时使用
type DepsInit struct {
}

// NewApp http.Server 支持多个 可分别支持 manage 和 portal API
func NewApp(logger log.Logger, conf *conf.Config, _ *DepsInit, gs *grpc.Server, hss ...*http.Server) *kratos.App {
	servers := []transport.Server{}
	for _, s := range hss {
		if s != nil {
			servers = append(servers, s)
		}
	}
	if gs != nil {
		servers = append(servers, gs)
	}
	if len(servers) == 0 {
		panic("请检查配置文件，至少保证启动 GRPC 或 HTTP 中的一个服务")
	}
	server := conf.Server
	return kratos.New(
		kratos.ID(server.Id),
		kratos.Name(server.Name),
		kratos.Version(server.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(servers...),
	)
}

// UpdateDefaultConfig 根据启动参数更新配置文件
func UpdateDefaultConfig(conf *conf.Config, id string, name string, version string) *conf.Config {
	if conf.Server != nil {
		if conf.Server.Id == "" {
			conf.Server.Id = id
		}
		if conf.Server.Name == "" {
			conf.Server.Name = name
		}
		if conf.Server.Version == "" {
			conf.Server.Version = version
		}
	}
	if conf.Trace != nil {
		if conf.Trace.Service == "" {
			conf.Trace.Service = name
		}
	}
	if conf.Metrics != nil {
		if conf.Metrics.Service == "" {
			conf.Metrics.Service = name
		}
	}

	return conf
}

// LoadFileConf 加载文件配置
func LoadFileConf(flagconf string) (config.Config, *conf.Config) {
	c := config.New(config.WithSource(file.NewSource(flagconf)))
	if err := c.Load(); err != nil {
		panic(err)
	}
	var bc conf.Config
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	return c, &bc
}
