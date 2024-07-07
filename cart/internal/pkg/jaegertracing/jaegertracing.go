package jaegertracing

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

type AddressExporter interface {
	GetTraceEndpointURL() string
	GetDeploymentEnvironment() string
}

// New - устанавливает и возвращает TracerProvider
func New(cfg AddressExporter, ServiceName string) (*trace.TracerProvider, error) {
	rootCtx := context.Background()

	exporter, err := otlptracehttp.New(rootCtx, otlptracehttp.WithEndpointURL(cfg.GetTraceEndpointURL()))
	if err != nil {
		return nil, err
	}

	res, err := resource.New(rootCtx,
		resource.WithAttributes(
			semconv.ServiceName(ServiceName),
			semconv.DeploymentEnvironment(cfg.GetTraceEndpointURL()),
		),
	)
	if err != nil {
		return nil, err
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(res),
	)

	otel.SetTracerProvider(tp)

	return tp, nil
}
