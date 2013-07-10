package main

import (
	"log"
	"github.com/nickstenning/router/triemux"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main () {
	mux := triemux.NewMux()

	googUrl, _ := url.Parse("http://google.com")
	aaplUrl, _ := url.Parse("http://apple.com")
	goog := httputil.NewSingleHostReverseProxy(googUrl)
	aapl := httputil.NewSingleHostReverseProxy(aaplUrl)

	// register a prefix route pointing to the Google backend (all requests to
	// "/google<anything>" will go to this backend)
	mux.Handle("/google", true, goog)

	// register an exact (non-prefix) route pointing to the Apple backend
	mux.Handle("/apple", false, aapl)

	log.Println("Listening on :8088")
	log.Fatalln(http.ListenAndServe(":8088", mux))
}
