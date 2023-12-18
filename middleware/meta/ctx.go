// Package meta ctx
package meta

import (
	"context"

	"github.com/go-kratos/kratos/v2/metadata"
)

// GetNewCtx _
func GetNewCtx(ctx context.Context) context.Context {
	newCtx := context.Background()
	md, ok := metadata.FromServerContext(ctx)
	if ok {
		newCtx = metadata.NewServerContext(newCtx, md)
	}
	return newCtx
}
