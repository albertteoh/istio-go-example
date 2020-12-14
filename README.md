# Istio Go Auto-Instrumentation Example

The purpose of this example is to explore low-to-no effort open-source instrumentation options available for distributed tracing and
understand tradeoffs between solutions.

This example consists of two simple Go microservices where `service-a` calls `service-b`. Both services expose a `/ping` endpoint,
that are instrumented using Istio.

For examples on manually instrumenting services with Jaeger, please see: [albertteoh/jaeger-go-example](https://github.com/albertteoh/jaeger-go-example).

The key takeaways are:
- There is still a need for code changes to [propagate context](https://istio.io/latest/faq/distributed-tracing/#how-to-support-tracing)
to allow correlating inbound with outbound calls.
- The code changes are fairly minimal and involve copying specific headers (Zipkin headers in the case of Istio) from inbound to outbound HTTP requests. "Leaf" services do not require any code changes, just services making outbound calls.
- One can choose to leverage the [OpenTracing API](https://opentracing.io/) and use an implementing tracer like [Jaeger](http://jaegertracing.io/docs/latest), or implement their own logic to do so.

Feedback and improvements (via PRs) are most welcome!

# Prerequisites

## Installing Istio & Jaeger
1. [Install Kubernetes](https://kubernetes.io/docs/setup/)

2. [Download](https://istio.io/latest/docs/setup/getting-started/#download) and
   [install](https://istio.io/latest/docs/setup/getting-started/#install) Istio.

3. [Install Jaeger Addon in Istio](https://istio.io/latest/docs/tasks/observability/distributed-tracing/jaeger/)

# Getting Started

## Start the example

Builds and starts the services in Istio.
```
$ make start
```

## Run the example

Hit `service-a`'s endpoint (via the istio-ingressgateway) to trigger the trace.
```
$ curl -w '\n' http://localhost:80/ping
```

## Validate

Should see `service-a -> service-b` on STDOUT.

The script will open Jaeger in a browser tab where you can select `service-a.default` from the "Service" dropdown and click the "Find Traces" button.

## Stop the example

Stop and remove containers.

```
$ make stop
```

# Additional References
- [Istio FAQ: How to support distributed tracing?](https://istio.io/latest/faq/distributed-tracing/#how-to-support-tracing)
- [Discussion on need for “hints” to enable correlation between inbound and outbound](https://discuss.istio.io/t/istio-tracing-and-correlation/2630)
- [Propagating Zipkin Trace Headers](https://github.com/jaegertracing/jaeger-client-go/blob/f7e0d4744fa6d5287c53b8ac8d4f83089ce07ce8/zipkin/README.md#L5)
