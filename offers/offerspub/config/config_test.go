package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "CheckENVSet",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Load()
			assert.Equal(t, os.Getenv("APP_NAME"), "offerspub", "Env must be set")
		})
	}
}
