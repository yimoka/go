// Package trace trace
package trace

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"
)

// WithReplyMiddleware 将 traceID 写入请求头
func WithReplyMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			reply, err = handler(ctx, req)
			if tr, ok := transport.FromServerContext(ctx); ok {
				x := tracing.TraceID()(ctx)
				traceID, _ := x.(string)
				tr.ReplyHeader().Set("X-Trace-Id", traceID)
			}
			return
		}
	}
}
