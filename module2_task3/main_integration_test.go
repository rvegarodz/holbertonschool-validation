package main

import (
	//nolint:staticcheck
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_server(t *testing.T) {
	if testing.Short() {
		t.Skip("Flag `-short` provided: skipping Integration Tests.")
	}

	tests := []struct {
		name         string
		URI          string
		responseCode int
		body         string
	}{
		{
			name:         "Home page",
			URI:          "",
			responseCode: 404,
			body:         "404 page not found\n",
		},
		{
			name:         "HealthCheck page",
			URI:          "/health",
			responseCode: 200,
			body:         "ALIVE",
		},
		{
			name:         "Hello page",
			URI:          "/hello?name=Holberton",
			responseCode: 200,
			body:         "Hello Holberton!",
		},
		{
			name:         "Grace Hopper",
			URI:          "/hello?name=Grace Hopper",
			responseCode: 200,
			body:         "Hello Grace Hopper!",
		},
		{
			name:         "Rosalind Franklin",
			URI:          "/hello?name=Rosalind Franklin",
			responseCode: 200,
			body:         "Hello Rosalind Franklin!",
		},
		{
			name:         "No name parameter",
			URI:          "/hello?",
			responseCode: 200,
			body:         "Hello there!",
		},
		{
			name:         "Empty name parameter",
			URI:          "/hello?name=",
			responseCode: 400,
			body:         "",
		},
		{
			name:         "Name with special characters",
			URI:          "/hello?name=John%20Doe",
			responseCode: 200,
			body:         "Hello John Doe!",
		},
		{
			name:         "Name with UTF-8 characters",
			URI:          "/hello?name=José",
			responseCode: 200,
			body:         "Hello José!",
		},
		{
			name:         "Multiple name parameters, only first one considered",
			URI:          "/hello?name=John&name=Jane",
			responseCode: 200,
			body:         "Hello John!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(setupRouter())
			defer ts.Close()

			res, err := http.Get(ts.URL + tt.URI)
			if err != nil {
				t.Fatal(err)
			}

			// Check that the status code is what you expect.
			expectedCode := tt.responseCode
			gotCode := res.StatusCode
			if gotCode != expectedCode {
				t.Errorf("handler returned wrong status code: got %q want %q", gotCode, expectedCode)
			}

			// Check that the response body is what you expect.
			expectedBody := tt.body
			bodyBytes, err := ioutil.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				t.Fatal(err)
			}
			gotBody := string(bodyBytes)
			if gotBody != expectedBody {
				t.Errorf("handler returned unexpected body: got %q want %q", gotBody, expectedBody)
			}
		})
	}
}
