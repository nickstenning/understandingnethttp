package main

import (
	"fmt"
	"log"
	"net/http"
)

const listenAddr = ":4000"

// START OMIT
func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello!")
}

func sayGoodbye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bye!")
}

func main() {
	log.Println("Listening on", listenAddr)

	http.HandleFunc("/hello", sayHello)
	http.HandleFunc("/bye", sayGoodbye)

	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
