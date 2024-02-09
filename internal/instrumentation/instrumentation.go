package instrumentation

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	"go.opentelemetry.io/otel/trace"
)

type TracerProvider interface {
	trace.TracerProvider
}

type SpanExporter interface {
	sdktrace.SpanExporter
}

type tracerProvider struct {
	trace.TracerProvider
}

func NewTracerProvider(exp SpanExporter, serviceName string) TracerProvider {
	res := resource.NewWithAttributes(semconv.SchemaURL, semconv.ServiceName(serviceName))
	tp := sdktrace.NewTracerProvider(sdktrace.WithBatcher(exp), sdktrace.WithResource(res))
	return &tracerProvider{tp}
}

func SetTraceProviderGlobally(tp TracerProvider) {
	otel.SetTracerProvider(tp)
}
