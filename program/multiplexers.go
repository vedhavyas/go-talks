package main

import (
	"log"
	"net/http"
)

type SimpleHandler struct{}

func (s SimpleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("simple handler "))
}

func main() {
	multiplexer := http.NewServeMux()
	multiplexer.HandleFunc("/func", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("simle handlefunc"))
	})
	multiplexer.Handle("/handler", SimpleHandler{})
	log.Fatal(http.ListenAndServe(":8080", multiplexer))
}
