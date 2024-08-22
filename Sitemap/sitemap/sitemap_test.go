package sitemap

import "testing"

// TestParseAndValidateURL tests the ParseAndValidateURL function for various valid and invalid URLs.
func TestParseAndValidateURL(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedError bool
		expectedURL   string
	}{
		{
			name:          "Valid URL with HTTP scheme",
			input:         "http://example.com",
			expectedError: false,
			expectedURL:   "http://example.com",
		},
		{
			name:          "Valid URL with HTTPS scheme",
			input:         "https://example.com",
			expectedError: false,
			expectedURL:   "https://example.com",
		},
		{
			name:          "URL without scheme, should default to HTTPS",
			input:         "example.com",
			expectedError: false,
			expectedURL:   "https://example.com",
		},
		{
			name:          "URL with subdomain and no scheme",
			input:         "sub.example.com",
			expectedError: false,
			expectedURL:   "https://sub.example.com",
		},
		{
			name:          "URL with scheme but no host",
			input:         "http://",
			expectedError: true,
			expectedURL:   "",
		},
		{
			name:          "Invalid URL without TLD",
			input:         "http://localhost",
			expectedError: true,
			expectedURL:   "",
		},
		{
			name:          "Invalid URL without any domain",
			input:         "https://",
			expectedError: true,
			expectedURL:   "",
		},
		{
			name:          "Completely invalid URL",
			input:         "this is not a url",
			expectedError: true,
			expectedURL:   "",
		},
		{
			name:          "Empty string",
			input:         "",
			expectedError: true,
			expectedURL:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseAndValidateURL(tt.input)
			if (err != nil) != tt.expectedError {
				t.Errorf("ParseAndValidateURL() error = %v, expectedError %v", err, tt.expectedError)
				return
			}
			if err == nil && got.String() != tt.expectedURL {
				t.Errorf("ParseAndValidateURL() = %v, expected %v", got.String(), tt.expectedURL)
			}
		})
	}
}
