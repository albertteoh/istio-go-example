package tracing

import (
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/zipkin"
)

// Init returns an instance of Jaeger Tracer.
func Init() (opentracing.Tracer, io.Closer) {
	// Zipkin shares span ID between client and server spans; it must be enabled via the following option.
	zipkinPropagator := zipkin.NewZipkinB3HTTPHeaderPropagator()

	// serviceName, sampler and reporter are not used since we're not emitting traces with this tracer.
	// This tracer is used for extracting and injecting trace header data, to propagate context.
	tracer, closer := jaeger.NewTracer(
		"",
		jaeger.NewConstSampler(false),
		jaeger.NewNullReporter(),
		jaeger.TracerOptions.Injector(opentracing.HTTPHeaders, zipkinPropagator),
		jaeger.TracerOptions.Extractor(opentracing.HTTPHeaders, zipkinPropagator),
	)

	opentracing.SetGlobalTracer(tracer)
	return tracer, closer
}
