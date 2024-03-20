// Package meta staff
package meta

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
)

// GetStaffID 获取员工 id
func GetStaffID(ctx context.Context) (string, *errors.Error) {
	return GetValue(ctx, staffIDKey)
}

// SetStaffID 设置用户 id
func SetStaffID(ctx context.Context, staffID string) context.Context {
	return SetValue(ctx, staffIDKey, staffID)
}
