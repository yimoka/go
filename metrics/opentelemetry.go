package metrics

import (
	"context"
	"time"

	"github.com/yimoka/go/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

var (
	MeterProviderName = ""
)

// SetMeterProvider 设置全局 MeterProvider
func SetMeterProvider(conf *config.Metrics) error {
	mp, err := NewMeterProvider(conf)
	if err != nil {
		return err
	}
	MeterProviderName = conf.Service
	otel.SetMeterProvider(mp)
	return nil
}

// NewMeterProvider 创建新的 MeterProvider
func NewMeterProvider(conf *config.Metrics) (*sdkmetric.MeterProvider, error) {
	timeout := conf.Timeout
	if timeout == 0 {
		timeout = 6
	}
	interval := conf.Interval
	if interval == 0 {
		interval = 20
	}
	opts := []otlpmetricgrpc.Option{
		otlpmetricgrpc.WithEndpoint(conf.Endpoint),
		otlpmetricgrpc.WithTimeout(time.Duration(timeout) * time.Second),
	}
	if conf.Insecure == nil || !*conf.Insecure {
		opts = append(opts, otlpmetricgrpc.WithInsecure())
	}
	if len(conf.Headers) > 0 {
		opts = append(opts, otlpmetricgrpc.WithHeaders(conf.Headers))
	}

	// 创建 OTLP gRPC metrics exporter
	metricsExporter, err := otlpmetricgrpc.New(context.Background(), opts...)
	if err != nil {
		return nil, err
	}
	// 构建资源属性
	attrs := []attribute.KeyValue{
		semconv.ServiceNameKey.String(conf.Service),
		attribute.String("env", conf.Env),
	}
	if conf.Namespace != "" {
		attrs = append(attrs, attribute.String("namespace", conf.Namespace))
	}
	if conf.Subsystem != "" {
		attrs = append(attrs, attribute.String("subsystem", conf.Subsystem))
	}

	// 添加自定义标签
	for k, v := range conf.Labels {
		attrs = append(attrs, attribute.String(k, v))
	}

	// 创建资源
	res, err := resource.New(context.Background(), resource.WithAttributes(attrs...))
	if err != nil {
		return nil, err
	}

	// 创建 MeterProvider
	provider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(metricsExporter, sdkmetric.WithInterval(time.Duration(interval)*time.Second))),
		sdkmetric.WithResource(res),
	)
	return provider, nil
}

// GetMeter 获取 Meter 实例
func GetMeter(conf *config.Metrics) metric.Meter {
	provider := otel.GetMeterProvider()
	if provider != nil {
		name := MeterProviderName
		if conf != nil && conf.Service != "" {
			name = conf.Service
		}
		return provider.Meter(name)
	}
	if conf == nil || conf.Endpoint == "" {
		panic("conf is nil or endpoint is empty")
	}
	provider, err := NewMeterProvider(conf)
	if err != nil {
		panic(err)
	}
	return provider.Meter(conf.Service)
}
