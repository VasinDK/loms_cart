package jaegertracing

import (
	"context"
	"route256/loms/internal/model"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	resource2 "go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
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

	resource, err := resource2.Merge(
		resource2.Default(),
		resource2.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(model.ServiceName),
			semconv.DeploymentEnvironment(cfg.GetTraceEndpointURL()),
		),
	)

	if err != nil {
		return nil, err
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return tp, nil
}
