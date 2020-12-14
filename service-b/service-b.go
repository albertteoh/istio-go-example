package main

import (
	"fmt"
	"log"
	"net/http"
)

const thisServiceName = "service-b"

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("%s", thisServiceName)))
	})
	log.Printf("Listening on localhost:8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
