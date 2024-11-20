package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock function to simulate a server endpoint
func mockServerResponse(w http.ResponseWriter, r *http.Request) {
	// Customize the response based on the requested URL, headers, etc.
	if r.URL.Path == "/api/v1/resource" {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"message": "mocked data response"}`)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, `{"error": "resource not found"}`)
}

// Function to create and return the mock server
func createMockServer() *httptest.Server {
	// Create a new HTTP server with the specified handler
	return httptest.NewServer(http.HandlerFunc(mockServerResponse))
}

// Test function to use the mock server
func TestApiClient(t *testing.T) {
	// Start the mock server and get its URL
	mockServer := createMockServer()
	defer mockServer.Close()

	// Mock server URL can be used in client calls
	apiUrl := mockServer.URL + "/api/v1/resource"
	resp, err := http.Get(apiUrl)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	expected := `{"message": "mocked data response"}` + "\n"
	if string(body) != expected {
		t.Errorf("Expected response %s, got %s", expected, body)
	}
}

// func main() {
// }
