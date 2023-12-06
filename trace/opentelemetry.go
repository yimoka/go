// Package trace opentelemetry
package trace

import (
	"context"

	"github.com/yimoka/go/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

// SetTracerProvider _
func SetTracerProvider(conf *config.Trace) error {
	exp, err := otlptracegrpc.New(context.Background(), otlptracegrpc.WithEndpoint(conf.Endpoint), otlptracegrpc.WithInsecure())
	if err != nil {
		return err
	}
	rate := 1.0
	if conf.SamplingRate > 0 {
		rate = float64(conf.SamplingRate)
	}

	attrs := []attribute.KeyValue{
		semconv.ServiceNameKey.String(conf.Service),
		attribute.String("env", conf.Env),
	}

	auth := conf.Auth
	if auth != nil {
		token := auth.Token
		if token != "" {
			attrs = append(attrs, attribute.String("token", token))
		} else {
			attrs = append(attrs, attribute.String("username", auth.Name))
			attrs = append(attrs, attribute.String("password", auth.Password))
		}
	}

	tp := tracesdk.NewTracerProvider(
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(rate))),
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewSchemaless(attrs...)),
	)
	otel.SetTracerProvider(tp)
	return nil
}
