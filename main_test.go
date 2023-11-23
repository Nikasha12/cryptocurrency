// main_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetPrice(t *testing.T) {
	// You can create mock responses using httptest.NewServer
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Mock the response from the external API for testing getPrice
		// For example, you can simulate a response for a specific cryptocurrency
		// Ensure to set the appropriate headers and status code
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"bitcoin": {"cad": 50000}}`))
	}))
	defer server.Close()

	// Replace the apiURL with the server URL in your implementation
	apiURL = server.URL + "/%s"

	price, err := getPrice("bitcoin")
	if err != nil {
		t.Errorf("getPrice returned an error: %v", err)
	}

	expectedPrice := 50000.0
	if price != expectedPrice {
		t.Errorf("Expected price: %.2f, got: %.2f", expectedPrice, price)
	}
}

func TestPriceHandler(t *testing.T) {
	// You can use httptest.NewRecorder to capture the HTTP response
	req, err := http.NewRequest("GET", "/price?crypto=bitcoin", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(priceHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Current price of bitcoin: "
	if !contains(rr.Body.String(), expected) {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestHTTPRequests(t *testing.T) {
	// You can use httptest.NewServer to mock the external API
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Mock the response from the external API for testing HTTP requests
		// For example, you can simulate a response for a specific cryptocurrency
		// Ensure to set the appropriate headers and status code
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"bitcoin": {"cad": 50000}}`))
	}))
	defer server.Close()

	// Replace the apiURL with the server URL in your implementation
	apiURL = server.URL + "/%s"

	// You can now test your functions that make HTTP requests
	// For example, calling getPrice or invoking the priceHandler through an HTTP request
}

// Add more tests as needed

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}
