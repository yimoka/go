// Package meta 适配 kratos 的元数据操作
package meta

import (
	"context"
	"strconv"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/yimoka/api/fault"
)

// GetValue 获取元数据
func GetValue(ctx context.Context, key string) (string, *errors.Error) {
	if md, ok := metadata.FromServerContext(ctx); ok {
		str := md.Get(globalPrefix + key)
		return str, nil
	}
	return "", fault.ErrorBadRequest("获取元数据失败")
}

// SetValue 设置元数据
func SetValue(ctx context.Context, kv ...string) context.Context {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		return ctx
	}
	for i, v := range kv {
		if i%2 != 0 {
			md.Set(globalPrefix+kv[i-1], v)
		}
	}
	return metadata.NewServerContext(ctx, md)
}

// GetIntValue 获取 int 类型的元数据
func GetIntValue(ctx context.Context, key string) (int, *errors.Error) {
	str, gErr := GetValue(ctx, key)
	if gErr != nil {
		return 0, gErr
	}
	if str == "" {
		return 0, nil
	}
	i, err := strconv.Atoi(str)
	if err != nil {
		rErr := fault.ErrorBadRequest("获取元数据 %s 转换为 int 失败", key)
		metadata := map[string]string{"source": key, "target": "int"}
		rErr.Metadata = metadata
		return 0, rErr
	}
	return i, nil
}

// SetIntValue 设置 int 类型的元数据
func SetIntValue(ctx context.Context, key string, value int) context.Context {
	return SetValue(ctx, key, strconv.Itoa(value))
}

// GetInt8Value 获取 int8 类型的元数据
func GetInt8Value(ctx context.Context, key string) (int8, *errors.Error) {
	str, gErr := GetValue(ctx, key)
	if gErr != nil {
		return 0, gErr
	}
	if str == "" {
		return 0, nil
	}
	i, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		rErr := fault.ErrorBadRequest("获取元数据 %s 转换为 int8 失败", key)
		metadata := map[string]string{"source": key, "target": "int8"}
		rErr.Metadata = metadata
		return 0, rErr
	}
	return int8(i), nil
}

// SetInt8Value 设置 int8 类型的元数据
func SetInt8Value(ctx context.Context, key string, value int8) context.Context {
	return SetValue(ctx, key, strconv.Itoa(int(value)))
}

// GetInt16Value 获取 int16 类型的元数据
func GetInt16Value(ctx context.Context, key string) (int16, *errors.Error) {
	str, gErr := GetValue(ctx, key)
	if gErr != nil {
		return 0, gErr
	}
	if str == "" {
		return 0, nil
	}
	i, err := strconv.ParseInt(str, 10, 16)
	if err != nil {
		rErr := fault.ErrorBadRequest("获取元数据 %s 转换为 int16 失败", key)
		metadata := map[string]string{"source": key, "target": "int16"}
		rErr.Metadata = metadata
		return 0, rErr
	}
	return int16(i), nil
}

// SetInt16Value 设置 int16 类型的元数据
func SetInt16Value(ctx context.Context, key string, value int16) context.Context {
	return SetValue(ctx, key, strconv.Itoa(int(value)))
}

// GetInt32Value 获取 int32 类型的元数据
func GetInt32Value(ctx context.Context, key string) (int32, *errors.Error) {
	str, gErr := GetValue(ctx, key)
	if gErr != nil {
		return 0, gErr
	}
	if str == "" {
		return 0, nil
	}
	i, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		rErr := fault.ErrorBadRequest("获取元数据 %s 转换为 int32 失败", key)
		metadata := map[string]string{"source": key, "target": "int32"}
		rErr.Metadata = metadata
		return 0, rErr
	}
	return int32(i), nil
}

// SetInt32Value 设置 int32 类型的元数据
func SetInt32Value(ctx context.Context, key string, value int32) context.Context {
	return SetValue(ctx, key, strconv.Itoa(int(value)))
}

// GetInt64Value 获取 int64 类型的元数据
func GetInt64Value(ctx context.Context, key string) (int64, *errors.Error) {
	str, gErr := GetValue(ctx, key)
	if gErr != nil {
		return 0, gErr
	}
	if str == "" {
		return 0, nil
	}
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		rErr := fault.ErrorBadRequest("获取元数据 %s 转换为 int64 失败", key)
		metadata := map[string]string{"source": key, "target": "int64"}
		rErr.Metadata = metadata
		return 0, rErr
	}
	return i, nil
}

// GetLocalValue 获取局部元数据
func GetLocalValue(ctx context.Context, key string) (string, *errors.Error) {
	if md, ok := metadata.FromServerContext(ctx); ok {
		str := md.Get(localPrefix + key)
		return str, nil
	}
	return "", fault.ErrorBadRequest("获取元数据失败")
}

// SetLocalValue 设置局部元数据
func SetLocalValue(ctx context.Context, kv ...string) context.Context {
	// 对 kv 中的 key 添加前缀 x-md-local-
	useKV := make([]string, len(kv))
	for i, v := range kv {
		if i%2 == 0 {
			useKV[i] = localPrefix + v
		} else {
			useKV[i] = v
		}
	}
	return metadata.AppendToClientContext(ctx, useKV...)
}

// GetLocalIntValue 获取局部 int 类型的元数据
func GetLocalIntValue(ctx context.Context, key string) (int, *errors.Error) {
	str, gErr := GetLocalValue(ctx, key)
	if gErr != nil {
		return 0, gErr
	}
	if str == "" {
		return 0, nil
	}
	i, err := strconv.Atoi(str)
	if err != nil {
		rErr := fault.ErrorBadRequest("获取元数据 %s 转换为 int 失败", key)
		metadata := map[string]string{"source": key, "target": "int"}
		rErr.Metadata = metadata
		return 0, rErr
	}
	return i, nil
}

// SetLocalIntValue 设置局部 int 类型的元数据
func SetLocalIntValue(ctx context.Context, key string, value int) context.Context {
	return SetLocalValue(ctx, key, strconv.Itoa(value))
}

// GetLocalInt8Value 获取局部 int8 类型的元数据
func GetLocalInt8Value(ctx context.Context, key string) (int8, *errors.Error) {
	str, gErr := GetLocalValue(ctx, key)
	if gErr != nil {
		return 0, gErr
	}
	if str == "" {
		return 0, nil
	}
	i, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		rErr := fault.ErrorBadRequest("获取元数据 %s 转换为 int8 失败", key)
		metadata := map[string]string{"source": key, "target": "int8"}
		rErr.Metadata = metadata
		return 0, rErr
	}
	return int8(i), nil
}

// SetLocalInt8Value 设置局部 int8 类型的元数据
func SetLocalInt8Value(ctx context.Context, key string, value int8) context.Context {
	return SetLocalValue(ctx, key, strconv.Itoa(int(value)))
}

// GetLocalInt16Value 获取局部 int16 类型的元数据
func GetLocalInt16Value(ctx context.Context, key string) (int16, *errors.Error) {
	str, gErr := GetLocalValue(ctx, key)
	if gErr != nil {
		return 0, gErr
	}
	if str == "" {
		return 0, nil
	}
	i, err := strconv.ParseInt(str, 10, 16)
	if err != nil {
		rErr := fault.ErrorBadRequest("获取元数据 %s 转换为 int16 失败", key)
		metadata := map[string]string{"source": key, "target": "int16"}
		rErr.Metadata = metadata
		return 0, rErr
	}
	return int16(i), nil
}

// SetLocalInt16Value 设置局部 int16 类型的元数据
func SetLocalInt16Value(ctx context.Context, key string, value int16) context.Context {
	return SetLocalValue(ctx, key, strconv.Itoa(int(value)))
}

// GetLocalInt32Value 获取局部 int32 类型的元数据
func GetLocalInt32Value(ctx context.Context, key string) (int32, *errors.Error) {
	str, gErr := GetLocalValue(ctx, key)
	if gErr != nil {
		return 0, gErr
	}
	if str == "" {
		return 0, nil
	}
	i, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		rErr := fault.ErrorBadRequest("获取元数据 %s 转换为 int32 失败", key)
		metadata := map[string]string{"source": key, "target": "int32"}
		rErr.Metadata = metadata
		return 0, rErr
	}
	return int32(i), nil
}

// SetLocalInt32Value 设置局部 int32 类型的元数据
func SetLocalInt32Value(ctx context.Context, key string, value int32) context.Context {
	return SetLocalValue(ctx, key, strconv.Itoa(int(value)))
}

// GetLocalInt64Value 获取局部 int64 类型的元数据
func GetLocalInt64Value(ctx context.Context, key string) (int64, *errors.Error) {
	str, gErr := GetLocalValue(ctx, key)
	if gErr != nil {
		return 0, gErr
	}
	if str == "" {
		return 0, nil
	}
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		rErr := fault.ErrorBadRequest("获取元数据 %s 转换为 int64 失败", key)
		metadata := map[string]string{"source": key, "target": "int64"}
		rErr.Metadata = metadata
		return 0, rErr
	}
	return i, nil
}

// SetLocalInt64Value 设置局部 int64 类型的元数据
func SetLocalInt64Value(ctx context.Context, key string, value int64) context.Context {
	return SetLocalValue(ctx, key, strconv.Itoa(int(value)))
}
