package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestMapHandler tests the MapHandler function.
func TestMapHandler(t *testing.T) {
	// Set up test cases
	tests := []struct {
		name               string
		path               string
		expectedStatusCode int
		expectedLocation   string
	}{
		{
			name:               "Path found, should redirect",
			path:               "/raz-go",
			expectedStatusCode: http.StatusFound,
			expectedLocation:   "https://github.com/Raz-y/Go-backOps-journey",
		},
		{
			name:               "Path not found, should fallback",
			path:               "/non-existent",
			expectedStatusCode: http.StatusOK,
			expectedLocation:   "",
		},
	}
	// Set up the URL paths and the fallback handler
	urlPaths := map[string]string{
		"/raz-backend": "https://github.com/Raz-y/backend-development-roadmap",
		"/raz-go":      "https://github.com/Raz-y/Go-backOps-journey",
	}

	fallbackhandler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("Fallback handler invoked"))
	})
	// Run the test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new request with the test case path
			req := httptest.NewRequest("GET", tt.path, nil)
			// Create a ResponseRecorder to capture the response
			rr := httptest.NewRecorder()

			// Get the handler and serve the HTTP request
			handler := MapHandler(urlPaths, fallbackhandler)
			handler.ServeHTTP(rr, req)

			// Validate the status code
			if rr.Code != tt.expectedStatusCode {
				t.Errorf("expected status code %d, got %d", tt.expectedStatusCode, rr.Code)
			}

			// Validate the Location header
			if tt.expectedStatusCode == http.StatusFound {
				location := rr.Header().Get("Location")
				if location != tt.expectedLocation {
					t.Errorf("expected redirect to %s, got %s", tt.expectedLocation, location)
				}
			} else if tt.expectedStatusCode == http.StatusOK {
				// Validate the fallback response
				expectedBody := "Fallback handler invoked"
				if rr.Body.String() != expectedBody {
					t.Errorf("expected fallback response %s, got %s", expectedBody, rr.Body.String())
				}
			}
		})
	}
}
