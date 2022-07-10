package dockersecretsenv

import (
	"os"
	"testing"
)

func TestLoadSecrets(t *testing.T) {

	tests := []struct {
		name    string
		content string
	}{
		{
			name:    "SECRET1",
			content: "content_secret_1",
		},
		{
			name:    "SECRET__2",
			content: "content_secret_2",
		},
	}
	LoadSecrets("test/run/secrets")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// expect to see the following environment variables:
			// expect var SECRET1
			if r := os.Getenv(tt.name); r != tt.content {
				t.Errorf("expected %s to be %s was %s", tt.name, tt.content, r)
			}

			// expect var SECRET_2
		})
	}
}
