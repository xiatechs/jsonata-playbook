package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonata(t *testing.T) {
	tests := []struct {
		jsonataCommand string
		input          string
		expectedOutput string
	}{
		{
			jsonataCommand: `[$]`,
			input:          `{"name": "Tam"}`,
			expectedOutput: "[\n {\n  \"name\": \"Tam\"\n }\n]",
		},
		{
			jsonataCommand: `/* a comment */ [$]`,
			input:          `{"name": "Tam"}`,
			expectedOutput: "[\n {\n  \"name\": \"Tam\"\n }\n]",
		},
		{
			jsonataCommand: `/* a comment and also a weird field */ $.{ "niceField": $."number #" }`,
			input:          `{"number #": "Tam"}`,
			expectedOutput: "{\n \"niceField\": \"Tam\"\n}",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.jsonataCommand, func(t *testing.T) {
			output := processJsonata(tt.input, tt.jsonataCommand)

			assert.Equal(t, tt.expectedOutput, output)
		})
	}
}
