package main

import (
	//nolint:staticcheck
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock HealthCheckHandler function
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ALIVE"))
}

// Mock HelloHandler function
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Name parameter missing", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello " + name + "!"))
}

func setupRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", HealthCheckHandler)
	mux.HandleFunc("/hello", HelloHandler)
	return mux
}

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
	// Test case for HealthCheckHandler
	t.Run("HealthCheckHandler", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/health", nil)
		rec := httptest.NewRecorder()

		HealthCheckHandler(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Errorf("HealthCheckHandler returned wrong status code: got %d, want %d", res.StatusCode, http.StatusOK)
		}

		expectedBody := "ALIVE"
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		gotBody := string(bodyBytes)
		if gotBody != expectedBody {
			t.Errorf("HealthCheckHandler returned unexpected body: got %q, want %q", gotBody, expectedBody)
		}
	})

	// Test case for HelloHandler
	t.Run("HelloHandler", func(t *testing.T) {
		name := "Holberton"
		req := httptest.NewRequest("GET", "/hello?name="+name, nil)
		rec := httptest.NewRecorder()

		HelloHandler(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Errorf("HelloHandler returned wrong status code: got %d, want %d", res.StatusCode, http.StatusOK)
		}

		expectedBody := "Hello " + name + "!"
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		gotBody := string(bodyBytes)
		if gotBody != expectedBody {
			t.Errorf("HelloHandler returned unexpected body: got %q, want %q", gotBody, expectedBody)
		}
	})
}
