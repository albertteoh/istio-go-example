package ping

import (
	"fmt"
	"net/http"
	libhttp "ping/lib/http"
	"ping/lib/tracing"
)

// Ping sends a ping request to the given hostPort, ensuring a new span is created
// for the downstream call, and associating the span to the parent span, if available
// in the provided context.
func Ping(origReq *http.Request, hostPort string) (string, error) {
	// Extract the trace headers as well as the requestID which Istio needs to correlate spans.
	reqID, spanCtx, err := tracing.Extract(origReq)
	if err != nil {
		return "", err
	}

	// Create the request.
	url := fmt.Sprintf("http://%s/ping", hostPort)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// Inject trace headers into the request in order for Istio to correlate outbound with inbound calls.
	if err := tracing.Inject(spanCtx, req, reqID); err != nil {
		return "", err
	}
	return libhttp.Do(req)
}
