package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/image/edit/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Image Edit"))
	})

	http.HandleFunc("/image/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// post images is not defined yet
			w.WriteHeader(404)
			w.Write([]byte("Not Found"))
			return
		}

		imageID := r.URL.Path[len("/image/"):]
		if imageID == "" {
			// list of images
			w.Write([]byte("Dummy list of images"))
			return
		}

		w.Write([]byte(imageID + " served!!"))
	})

	// handle rest
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Unknown request"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
