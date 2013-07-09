package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

const listenAddr = ":4000"

func main() {
	log.Println("Listening on", listenAddr)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // HLhandler
		fmt.Fprintf(w, "Hello, %v", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(listenAddr, nil)) // HLhandler
}
