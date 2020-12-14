package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"ping/lib/ping"
	"ping/lib/tracing"

	"github.com/opentracing/opentracing-go"
)

const thisServiceName = "service-a"

func main() {
	tracer, closer := tracing.Init(thisServiceName)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	outboundHostPort, ok := os.LookupEnv("OUTBOUND_HOST_PORT")
	if !ok {
		outboundHostPort = "localhost:8082"
	}

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		response, err := ping.Ping(r, outboundHostPort)
		if err != nil {
			log.Fatalf("Error occurred: %s", err)
		}
		w.Write([]byte(fmt.Sprintf("%s -> %s", thisServiceName, response)))
	})
	log.Printf("Listening on localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
