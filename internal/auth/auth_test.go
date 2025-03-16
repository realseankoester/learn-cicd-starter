package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	testCases := []struct {
		name        string
		headers     http.Header
		expectKey   string
		expectError bool
	}{
		{
			name: "valid API key",
			headers: http.Header{
				"Authorization": []string{"ApiKey valid-api-key"},
			},
			expectKey:   "valid-api-key",
			expectError: false,
		},
		{
			name:        "missing API key",
			headers:     http.Header{},
			expectKey:   "",
			expectError: true,
		},
		{
			name: "malformed authorization header",
			headers: http.Header{
				"Authorization": []string{"malformed-header"},
			},
			expectKey:   "",
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/", nil)
			req.Header = tc.headers

			key, err := GetAPIKey(req.Header)

			if tc.expectError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tc.expectError && err != nil {
				t.Errorf("expected no error but got: %v", err)
			}

			if key != tc.expectKey {
				t.Errorf("expected key %q but got %q", tc.expectKey, key)
			}
		})
	}
}
