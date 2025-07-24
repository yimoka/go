package metrics

import (
	"context"

	"github.com/yimoka/go/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

// SetMeterProvider 设置全局 MeterProvider
func SetMeterProvider(conf *config.Metrics) error {
	mp, err := NewMeterProvider(conf)
	if err != nil {
		return err
	}
	otel.SetMeterProvider(mp)
	return nil
}

// NewMeterProvider 创建新的 MeterProvider
func NewMeterProvider(conf *config.Metrics) (*sdkmetric.MeterProvider, error) {
	// 创建 Prometheus exporter
	exporter, err := prometheus.New()
	if err != nil {
		return nil, err
	}

	// 构建资源属性
	attrs := []attribute.KeyValue{
		semconv.ServiceNameKey.String(conf.Service),
		attribute.String("env", conf.Env),
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
		sdkmetric.WithReader(exporter),
		sdkmetric.WithResource(res),
	)

	return provider, nil
}

// GetMeter 获取 Meter 实例
func GetMeter(conf *config.Metrics) metric.Meter {
	provider := otel.GetMeterProvider()
	if provider != nil {
		return provider.Meter(conf.Service)
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
