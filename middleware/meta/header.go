// Package meta header
package meta

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport"
)

// GetRequestHeader 获取相关
func GetRequestHeader(ctx context.Context) (transport.Header, bool) {
	t, boo := transport.FromServerContext(ctx)
	if !boo {
		return nil, boo
	}
	return t.RequestHeader(), boo
}

// GetRequestHeaderVal 获取 Header 的 val 值
func GetRequestHeaderVal(ctx context.Context, key string) string {
	h, b := GetRequestHeader(ctx)
	if !b {
		return ""
	}
	return h.Get(key)
}

// GetAuthorization 获取 Header Authorization 的值
func GetAuthorization(ctx context.Context) string {
	return GetRequestHeaderVal(ctx, "Authorization")
}

// GetUserAgent 获取 Header User-Agent 的值
func GetUserAgent(ctx context.Context) string {
	return GetRequestHeaderVal(ctx, "User-Agent")
}

// GetContentType 获取 Header Content-Type 的值
func GetContentType(ctx context.Context) string {
	return GetRequestHeaderVal(ctx, "Content-Type")
}

// GetAcceptLanguage 获取 Header Accept-Language 的值
func GetAcceptLanguage(ctx context.Context) string {
	return GetRequestHeaderVal(ctx, "Accept-Language")
}

// GetReplyHeader _
func GetReplyHeader(ctx context.Context, key, value string) (transport.Header, bool) {
	t, ok := transport.FromServerContext(ctx)
	if !ok {
		return nil, ok
	}
	return t.ReplyHeader(), ok
}
