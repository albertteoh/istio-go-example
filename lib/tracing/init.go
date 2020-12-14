package tracing

import (
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/zipkin"
)

// Init returns an instance of Jaeger Tracer.
func Init(service string) (opentracing.Tracer, io.Closer) {
	// Zipkin shares span ID between client and server spans; it must be enabled via the following option.
	zipkinPropagator := zipkin.NewZipkinB3HTTPHeaderPropagator()
	tracer, closer := jaeger.NewTracer(
		service,
		jaeger.NewConstSampler(false),
		jaeger.NewNullReporter(),
		jaeger.TracerOptions.Injector(opentracing.HTTPHeaders, zipkinPropagator),
		jaeger.TracerOptions.Extractor(opentracing.HTTPHeaders, zipkinPropagator),
		jaeger.TracerOptions.ZipkinSharedRPCSpan(true),
	)

	return tracer, closer
}
