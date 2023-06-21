package httpclient

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHttpClient_Get(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello world!"))
	}))

	defer server.Close()

	timeout := 5 * time.Second
	client := New(timeout)

	response, err := client.Get(server.URL)
	if err != nil {
		t.Errorf("Error making GET request: %v", err)
	}

	expectedResponse := "Hello world!"
	if string(response) != expectedResponse {
		t.Errorf("Unexpected response. Expected: %s, Got: %s", expectedResponse, string(response))
	}

}
