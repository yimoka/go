// Package logger  logger.go
package logger

import (
	"os"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/yimoka/go/config"
)

// GetLogger _
func GetLogger(conf *config.Config) log.Logger {
	logger := conf.Logger
	if logger == nil {
		return GetStdLogger(conf)
	}
	if logger.Provider == "tencent" {
		return GetTencentLogger(conf)
	}
	return GetStdLogger(conf)
}

// GetStdLogger _
func GetStdLogger(conf *config.Config) log.Logger {
	logger := log.NewStdLogger(os.Stdout)
	return getLogger(conf.Server, logger)
}

// GetLogger _
func getLogger(service *config.Server, logger log.Logger) log.Logger {
	return log.With(logger,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", service.Id,
		"service.name", service.Name,
		"service.version", service.Version,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
}
