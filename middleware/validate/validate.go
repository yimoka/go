// Package validate 提供用于验证协议缓冲区消息的中间件
package validate

import (
	"context"

	"github.com/bufbuild/protovalidate-go"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/yimoka/api/fault"
	"github.com/yimoka/go/lang"
	"google.golang.org/protobuf/proto"
)

// Option 是验证中间件的配置选项函数类型
type Option func(*options)

// options 包含验证中间件的配置选项
type options struct {
	// validateMsg 是自定义的验证错误消息处理函数
	// 如果提供此函数，将优先使用它来处理验证错误消息
	// ctx: 上下文
	// violation: 验证违规信息
	// langs: 语言选项
	// 返回值:
	// - msg: 错误消息
	// - isMatch: 是否匹配到自定义错误消息
	validateMsg func(ctx context.Context, violation *protovalidate.Violation, langs ...string) (msg string, isMatch bool)

	// commonLang 用于处理通用的多语言消息
	commonLang *lang.CommonLang
}

// WithCommonLang 设置通用语言处理器
func WithCommonLang(l *lang.CommonLang) Option {
	return func(o *options) {
		o.commonLang = l
	}
}

// WithValidateMsg 设置自定义验证错误消息处理函数
func WithValidateMsg(validateMsg func(ctx context.Context, violation *protovalidate.Violation, langs ...string) (msg string, isMatch bool)) Option {
	return func(o *options) {
		o.validateMsg = validateMsg
	}
}

// getDefaultErrorMsg 获取验证错误的默认消息
// 优先级：
// 1. Proto.Message - 包含人类可读的错误消息
// 2. FieldValue - 包含导致验证失败的字段值
// 3. 默认消息
func getDefaultErrorMsg(v *protovalidate.Violation) string {
	if v.Proto != nil && v.Proto.Message != nil {
		return *v.Proto.Message
	}

	// 如果有字段值，添加到错误消息中
	if v.FieldValue.IsValid() {
		return "invalid value: " + v.FieldValue.String()
	}

	return "validation failed"
}

// ProtoValidate 返回一个用于协议缓冲区消息验证的中间件
func ProtoValidate(opts ...Option) middleware.Middleware {
	options := &options{}
	for _, opt := range opts {
		opt(options)
	}

	validator, err := protovalidate.New()
	if err != nil {
		panic(err)
	}

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			// 检查请求是否为 protobuf 消息
			if msg, ok := req.(proto.Message); ok {
				if err := validator.Validate(msg); err != nil {
					var valErr *protovalidate.ValidationError

					// 获取参数错误的基础消息
					parameterErrorMsg := "parameter error"
					if options.commonLang != nil {
						parameterErrorMsg = options.commonLang.GetParameterErrorMsg(ctx)
					}

					// 如果不是验证错误，返回基础错误信息
					if ok := errors.As(err, &valErr); !ok {
						return nil, fault.ErrorBadRequest("%s: %s", parameterErrorMsg, err.Error())
					}

					// 处理每个验证违规
					metadata := make(map[string]string)
					for _, v := range valErr.Violations {
						field := v.Proto.Field.String()
						constraintID := v.Proto.GetConstraintId()
						// 如果约束ID为空，则使用默认的错误消息
						if constraintID == "" {
							metadata[field] = getDefaultErrorMsg(v)
							continue
						}

						// 1. 尝试使用自定义错误消息处理
						if options.validateMsg != nil {
							if msg, isMatch := options.validateMsg(ctx, v); isMatch {
								metadata[field] = msg
								continue
							}
						}

						// 2. 尝试使用通用语言处理器的错误消息
						if options.commonLang != nil {
							if msg, isMatch := options.commonLang.GetValidateErrorMsg(ctx, v); isMatch {
								metadata[field] = msg
								continue
							}
						}

						metadata[field] = getDefaultErrorMsg(v)
					}

					// 构建错误响应
					err := fault.ErrorBadRequest("%s", parameterErrorMsg)
					if len(metadata) > 0 {
						err = err.WithMetadata(metadata)
					}
					return nil, err
				}
			}
			return handler(ctx, req)
		}
	}
}
