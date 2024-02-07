package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// func TestHello(t *testing.T) {
// 	got := Hello()
// 	want := "hello backend"

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

func TestGetPosts(t *testing.T) {
	// Create a new request to our API endpoint
	req, err := http.NewRequest("GET", "/posts", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a handler using the main() function
	handler := http.HandlerFunc(getPostsHandler)

	// Serve the HTTP request to the ResponseRecorder
	handler.ServeHTTP(rr, req)

	// Check the status code returned by the handler
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the content type header
	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned wrong content type: got %v want %v", contentType, expectedContentType)
	}
}
