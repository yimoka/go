// Package meta user
package meta

import "context"

const userIDKey = globalPrefix + "user-id"

// GetUserID 获取用户 id
func GetUserID(ctx context.Context) (string, error) {
	return GetValue(ctx, userIDKey)
}

// SetUserID 设置用户 id
func SetUserID(ctx context.Context, userID string) context.Context {
	return SetValue(ctx, userIDKey, userID)
}
