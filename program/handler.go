package main

import (
	"fmt"
	"log"
	"net/http"
)

type CountHandler struct {
	Count int
}

func (h *CountHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Count += 1
	responseString := fmt.Sprintf("You have hit %s and count is %v", r.URL.Path, h.Count)
	w.Write([]byte(responseString))
}

func main() {
	http.Handle("/count", &CountHandler{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
