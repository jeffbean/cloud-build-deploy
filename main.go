package main

import (
	"flag"
	"fmt"
	"net/http"
)

var port = flag.Int("port", 8080, "Port number to serve default backend 404 page.")

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Testing")
	})
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
