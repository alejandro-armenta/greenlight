package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":9000", "Server address")

	flag.Parse()

	log.Printf("starting server on %s", *addr)

	http.ListenAndServe(*addr, http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ale"))
		}))
}
