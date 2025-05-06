package auth_test

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetApiKey(t *testing.T) {
	tt := []struct {
		apiKey      string
		parsedKey   string
		shouldError bool
	}{
		{
			apiKey:      "someapikey",
			parsedKey:   "",
			shouldError: true,
		},
		{
			apiKey:      "ApiKey someotherapikey",
			parsedKey:   "someotherapikey",
			shouldError: false,
		},
	}

	for _, tc := range tt {
		headers := http.Header{}
		headers.Add("Authorization", tc.apiKey)
		apiKey, err := auth.GetAPIKey(headers)
		if tc.shouldError && err == nil {
			t.Fatalf("did not receive expected error. Expected: %v, Got: %v", tc.shouldError, err)
		}

		if !tc.shouldError {
			if apiKey != tc.parsedKey {
				t.Fatalf("api keys do not match. Expected: %v, Got: %v", tc.parsedKey, apiKey)
			}
		}
	}
}
