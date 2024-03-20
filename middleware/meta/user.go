// Package meta user
package meta

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
)

// GetUserID 获取用户 id
func GetUserID(ctx context.Context) (string, *errors.Error) {
	return GetValue(ctx, userIDKey)
}

// SetUserID 设置用户 id
func SetUserID(ctx context.Context, userID string) context.Context {
	return SetValue(ctx, userIDKey, userID)
}
