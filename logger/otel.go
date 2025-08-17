// Package logger otel.go
package logger

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/yimoka/go/config"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	otellog "go.opentelemetry.io/otel/log"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GetOtelLogger _
func GetOtelLogger(conf *config.Config) log.Logger {
	otelLogger, err := NewOtelLogger(conf, conf.Server)
	if err != nil {
		panic(fmt.Sprintf("init otel logger error: %v", err))
	}

	return getLogger(conf.Server, otelLogger, conf.Logger)
}

// OtelLogger _
type OtelLogger interface {
	log.Logger
	Close() error
}

type otelLog struct {
	loggerProvider *sdklog.LoggerProvider
	resource       *resource.Resource
	opts           *config.Logger
	ctx            context.Context
	stdLogger      log.Logger
}

func (l *otelLog) Close() error {
	var errs []error
	if l.loggerProvider != nil {
		if err := l.loggerProvider.Shutdown(context.Background()); err != nil {
			errs = append(errs, fmt.Errorf("shutdown logger provider: %w", err))
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("errors during shutdown: %v", errs)
	}
	return nil
}

func (l *otelLog) Log(level log.Level, keyValues ...interface{}) error {
	// 创建日志记录
	record := otellog.Record{}
	record.SetTimestamp(time.Now())
	record.SetSeverity(convertLogLevel(level))
	record.SetBody(otellog.StringValue(extractMessage(keyValues)))

	// 添加属性
	attrs := make([]otellog.KeyValue, 0, len(keyValues)/2)
	for i := 0; i < len(keyValues); i += 2 {
		if i+1 < len(keyValues) {
			key := toString(keyValues[i])
			value := toString(keyValues[i+1])
			attrs = append(attrs, otellog.String(key, value))
		}
	}
	record.AddAttributes(attrs...)

	// 发送日志记录
	if l.loggerProvider != nil {
		logger := l.loggerProvider.Logger("logger")
		if l.ctx == nil {
			logger.Emit(context.Background(), record)
		} else {
			logger.Emit(l.ctx, record)
		}
		if l.stdLogger != nil {
			_ = l.stdLogger.Log(level, keyValues...)
		}
	}

	return nil
}

// convertLogLevel 将 go-kratos 日志级别转换为 OpenTelemetry 日志级别
func convertLogLevel(level log.Level) otellog.Severity {
	switch level {
	case log.LevelDebug:
		return otellog.SeverityDebug
	case log.LevelInfo:
		return otellog.SeverityInfo
	case log.LevelWarn:
		return otellog.SeverityWarn
	case log.LevelError, log.LevelFatal:
		return otellog.SeverityError
	default:
		return otellog.SeverityInfo
	}
}

// extractMessage 从键值对中提取消息
func extractMessage(keyValues []interface{}) string {
	for i := 0; i < len(keyValues); i += 2 {
		if i+1 < len(keyValues) {
			key := toString(keyValues[i])
			if key == "msg" || key == "message" || key == "operation" {
				return toString(keyValues[i+1])
			}
		}
	}
	return "日志记录"
}

// NewOtelLogger _
func NewOtelLogger(conf *config.Config, serverConfig *config.Server) (OtelLogger, error) {
	var loggerProvider *sdklog.LoggerProvider
	var res *resource.Resource
	var err error

	// 创建资源，设置服务信息
	res, err = resource.New(context.Background(),
		resource.WithAttributes(
			semconv.ServiceName(serverConfig.Name),
			semconv.ServiceVersion(serverConfig.Version),
			semconv.ServiceInstanceID(serverConfig.Id),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	opts := conf.Logger
	// 获取 endpoint
	endpoint := opts.Endpoint
	if endpoint == "" {
		return nil, fmt.Errorf("Logger.Endpoint must be specified")
	}

	// 准备 OTLP 日志导出器选项
	otlpOptions := []otlploggrpc.Option{
		otlploggrpc.WithEndpoint(endpoint),
		otlploggrpc.WithDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
	}
	if opts.Insecure == nil || !*opts.Insecure {
		otlpOptions = append(otlpOptions, otlploggrpc.WithInsecure())
	}
	if len(opts.Headers) > 0 {
		otlpOptions = append(otlpOptions, otlploggrpc.WithHeaders(opts.Headers))
	}

	// 添加认证配置
	if opts.Token != "" {
		// 使用 token 认证
		headers := map[string]string{
			"Authorization": "Bearer " + opts.Token,
		}
		otlpOptions = append(otlpOptions, otlploggrpc.WithHeaders(headers))
	} else if opts.Username != "" && opts.Password != "" {
		// 使用用户名密码认证
		headers := map[string]string{
			"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(opts.Username+":"+opts.Password)),
		}
		otlpOptions = append(otlpOptions, otlploggrpc.WithHeaders(headers))
	}

	// 创建日志导出器
	logExporter, err := otlploggrpc.New(
		context.Background(),
		otlpOptions...,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create OTLP log exporter: %w", err)
	}

	// 创建日志处理器
	processor := sdklog.NewBatchProcessor(logExporter)

	// 创建日志提供者
	loggerProvider = sdklog.NewLoggerProvider(
		sdklog.WithProcessor(processor),
		sdklog.WithResource(res),
	)

	otelLog := &otelLog{
		loggerProvider: loggerProvider,
		resource:       res,
		opts:           opts,
	}
	if opts.AlsoStd {
		stdLogger := GetStdLogger(conf)
		otelLog.stdLogger = stdLogger
	}
	return otelLog, nil
}
