// Package server grpc.go
package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/yimoka/go/config"
	"go.opentelemetry.io/otel"
)

// CreateGRPCServer new an gRPC server.
func CreateGRPCServer(conf *config.ServerItem, logger log.Logger, ms ...middleware.Middleware) *grpc.Server {
	if conf == nil || conf.Addr == "" {
		return nil
	}

	var opts = []grpc.ServerOption{}
	if conf.Network != "" {
		opts = append(opts, grpc.Network(conf.Network))
	}

	if conf.Addr != "" {
		opts = append(opts, grpc.Address(conf.Addr))
	}

	if conf.Timeout != nil {
		opts = append(opts, grpc.Timeout(conf.Timeout.AsDuration()))
	}

	use := []middleware.Middleware{}
	if conf.IsLog {
		use = append(use, logging.Server(logger))
	}

	if conf.IsTrace {
		use = append(use, tracing.Server(tracing.WithTracerProvider(otel.GetTracerProvider())))
	}

	use = append(use, metadata.Server(), validate.Validator())
	if len(ms) > 0 {
		use = append(use, ms...)
	}

	use = append(use, recovery.Recovery())

	opts = append(opts, grpc.Middleware(use...))

	srv := grpc.NewServer(opts...)

	return srv
}
