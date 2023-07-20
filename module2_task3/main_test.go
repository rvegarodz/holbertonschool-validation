package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func Test_HelloHandler(t *testing.T) {
	if !testing.Short() {
		t.Skip("Flag `-short` absent: skipping Unit Tests.")
	}

	tests := []struct {
		name         string
		queryString  string
		responseCode int
		body         string
	}{
		{
			name:         "Grace Hopper",
			queryString:  "name=Grace Hopper",
			responseCode: 200,
			body:         "Hello Grace Hopper!",
		},
		{
			name:         "Rosalind Franklin",
			queryString:  "name=Rosalind Franklin",
			responseCode: 200,
			body:         "Hello Rosalind Franklin!",
		},
		// INSERT MORE TESTS HERE
		{
			name:         "No name parameter",
			queryString:  "",
			responseCode: 200,
			body:         "Hello there!",
		},
		{
			name:         "Empty name parameter",
			queryString:  "name=",
			responseCode: 400,
			body:         "",
		},
		{
			name:         "Name with special characters",
			queryString:  "name=John%20Doe",
			responseCode: 200,
			body:         "Hello John Doe!",
		},
		{
			name:         "Name with UTF-8 characters",
			queryString:  "name=José",
			responseCode: 200,
			body:         "Hello José!",
		},
		{
			name:         "Multiple name parameters, only first one considered",
			queryString:  "name=John&name=Jane",
			responseCode: 200,
			body:         "Hello John!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Build an HTTP request for the current test `tt`
			req, err := http.NewRequest("GET", "/hello", nil)
			if err != nil {
				t.Fatal(err)
			}

			queryParams, _ := url.ParseQuery(tt.queryString)
			req.URL.RawQuery = queryParams.Encode()

			// Initialize a new httptest serve with the function under test
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(HelloHandler)
			handler.ServeHTTP(rr, req)

			// Check that the status code is what you expect.
			expectedCode := tt.responseCode
			gotCode := rr.Code
			if gotCode != expectedCode {
				t.Errorf("handler returned wrong status code: got %v want %v", gotCode, expectedCode)
			}

			// Check that the response body is what you expect.
			expectedBody := tt.body
			gotBody := rr.Body.String()
			if gotBody != expectedBody {
				t.Errorf("handler returned unexpected body: got %v want %v", gotBody, expectedBody)
			}
		})
	}
}
