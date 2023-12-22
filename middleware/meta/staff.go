// Package meta staff
package meta

import "context"

const staffIDKey = globalPrefix + "staff-id"

// GetStaffID 获取员工 id
func GetStaffID(ctx context.Context) (string, error) {
	return GetValue(ctx, staffIDKey)
}

// SetStaffID 设置用户 id
func SetStaffID(ctx context.Context, staffID string) context.Context {
	return SetValue(ctx, staffIDKey, staffID)
}
