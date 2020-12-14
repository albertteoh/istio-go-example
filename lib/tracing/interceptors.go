package tracing

import (
	"net/http"

	"github.com/opentracing/opentracing-go"
)

// Inject injects the outbound HTTP request with the given span's context to ensure
// correct propagation of span context throughout the trace.
func Inject(spanContext opentracing.SpanContext, request *http.Request, requestID string) error {
	request.Header.Add("x-request-id", requestID)
	return opentracing.GlobalTracer().Inject(
		spanContext,
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(request.Header))
}

// Extract extracts the inbound HTTP request to obtain the parent span's context to ensure
// correct propagation of span context throughout the trace.
func Extract(r *http.Request) (string, opentracing.SpanContext, error) {
	requestID := r.Header.Get("x-request-id")
	spanCtx, err := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(r.Header))
	return requestID, spanCtx, err
}
