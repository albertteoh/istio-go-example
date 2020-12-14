# Istio Go Auto-Instrumentation Example
Two simple Go microservices where `service-a` calls `service-b`. Both services expose a `/ping` endpoint,
that are "auto" instrumented with Istio.

For examples on manually instrumenting services with Jaeger, please see: [albertteoh/jaeger-go-example](https://github.com/albertteoh/jaeger-go-example).

The main code change required to enable tracing instrumentation with Istio is to propagate trace headers within each service to enable correlating inbound with outbound calls.

# Prerequisites

## Installing Istio & Jaeger
1. [Download](https://istio.io/latest/docs/setup/getting-started/#download) and
   [install](https://istio.io/latest/docs/setup/getting-started/#install) Istio.

2. [Install Jaeger Addon in Istio](https://istio.io/latest/docs/tasks/observability/distributed-tracing/jaeger/)

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
