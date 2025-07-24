// Package server http.go
package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/yimoka/go/config"
	ymetrics "github.com/yimoka/go/middleware/metrics"
	"github.com/yimoka/go/middleware/trace"
)

// CreateHTTPServer new an HTTP server.
func CreateHTTPServer(conf *config.ServerItem, traceConf *config.Trace, logger log.Logger, ms ...middleware.Middleware) *http.Server {
	if conf == nil || conf.Addr == "" {
		return nil
	}
	opts := []http.ServerOption{}
	opts = append(opts, http.Address(conf.Addr))
	if conf.Network != "" {
		opts = append(opts, http.Network(conf.Network))
	}

	if conf.Timeout != nil {
		opts = append(opts, http.Timeout(conf.Timeout.AsDuration()))
	}

	use := []middleware.Middleware{}

	use = append(use, recovery.Recovery())

	if conf.IsTrace {
		use = append(use, trace.CreateMiddleware(conf, traceConf)...)
	}

	if conf.IsLog {
		use = append(use, logging.Server(logger))
	}

	if conf.IsMetrics {
		metricRequests, metricSeconds := ymetrics.GetDefaultMetrics(nil)
		use = append(use, metrics.Server(
			metrics.WithSeconds(metricSeconds),
			metrics.WithRequests(metricRequests),
		))
	}

	use = append(use, metadata.Server())

	if len(ms) > 0 {
		use = append(use, ms...)
	}

	opts = append(opts, http.Middleware(use...))

	return http.NewServer(opts...)
}
