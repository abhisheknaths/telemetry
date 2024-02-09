package instrumentation

import (
	"context"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
)

func NewHTTPTraceExporter(ctx context.Context, endpoint, path string) (SpanExporter, error) {
	endpointOption := otlptracehttp.WithEndpoint(endpoint)
	securityOption := otlptracehttp.WithInsecure()
	urlPathOption := otlptracehttp.WithURLPath(path)
	exporter, err := otlptracehttp.New(ctx, endpointOption, securityOption, urlPathOption)
	if err != nil {
		return nil, err
	}
	return exporter, nil
}
