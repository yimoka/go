// Package server http.go
package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/yimoka/go/config"
	"github.com/yimoka/go/middleware/trace"
)

// CreateHTTPServer new an HTTP server.
func CreateHTTPServer(conf *config.ServerItem, traceConf *config.Trace, logger log.Logger, ms ...middleware.Middleware) *http.Server {
	if conf == nil || conf.Addr == "" {
		return nil
	}
	opts := []http.ServerOption{}
	use := []middleware.Middleware{}
	opts = append(opts, http.Address(conf.Addr))
	if conf.Network != "" {
		opts = append(opts, http.Network(conf.Network))
	}

	if conf.Timeout != nil {
		opts = append(opts, http.Timeout(conf.Timeout.AsDuration()))
	}

	if conf.IsLog {
		use = append(use, logging.Server(logger))
	}

	use = append(use, trace.CreateMiddleware(conf, traceConf)...)

	use = append(use, metadata.Server(), validate.Validator())

	if len(ms) > 0 {
		use = append(use, ms...)
	}

	use = append(use, recovery.Recovery())
	opts = append(opts, http.Middleware(use...))

	return http.NewServer(opts...)
}
