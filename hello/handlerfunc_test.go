package hello

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	request, _ := http.NewRequest("GET", "/hello", nil)

	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	if recorder.Body.String() != "Hello, Dev!!" {
		t.Fail()
	}
}
