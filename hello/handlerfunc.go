package hello

import "net/http"

// HelloHandler has a same method signature of handlerfunc.
// Hence this can be passed a handlerfunc to a specific url pattern
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Dev!!"))
}
