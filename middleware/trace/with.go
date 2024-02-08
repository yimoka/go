// Package trace trace
package trace

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/samber/lo"
	"github.com/yimoka/go/config"
	"github.com/yimoka/go/trace"
	"go.opentelemetry.io/otel"
)

// WithReplyMiddleware 将 traceID 写入请求头
func WithReplyMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			reply, err = handler(ctx, req)
			if tr, ok := transport.FromServerContext(ctx); ok {
				x := tracing.TraceID()(ctx)
				traceID, _ := x.(string)
				if traceID != "" {
					tr.ReplyHeader().Set("X-Trace-Id", traceID)
				}
			}
			return
		}
	}
}

// CreateMiddleware 创建 trace 中间件
func CreateMiddleware(conf *config.ServerItem, traceConfig *config.Trace) []middleware.Middleware {
	if conf == nil || conf.Addr == "" || !conf.IsTrace || traceConfig == nil {
		return nil
	}
	prefix := conf.TracePrefix
	suffix := conf.TraceSuffix
	if prefix != "" || suffix != "" {
		newService := prefix + traceConfig.Service + suffix
		newTrace := lo.FromPtr(traceConfig)
		newTrace.Service = newService
		provider, err := trace.NewTracerProvider(&newTrace)
		if err != nil {
			panic(err)
		}
		return []middleware.Middleware{tracing.Server(tracing.WithTracerProvider(provider)), WithReplyMiddleware()}
	}
	return []middleware.Middleware{tracing.Server(tracing.WithTracerProvider(otel.GetTracerProvider())), WithReplyMiddleware()}
}
