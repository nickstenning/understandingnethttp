package main

import (
	"fmt"
	"log"
	"net/http"
)

const listenAddr = ":4000"

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func main() {
	log.Println("Listening on", listenAddr)

	http.HandleFunc("/", sayHello) // HL
	http.ListenAndServe(listenAddr, nil) // HL
}
