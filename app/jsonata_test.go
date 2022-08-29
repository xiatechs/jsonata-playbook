package app

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestHandler(t *testing.T) {
	tests := []struct {
		method         string
		name           string
		testInput      string
		expectedOutput string
	}{
		{
			method: http.MethodPost,
			name:   `valid input`,
			testInput: `{
				"input": "{\"name\": \"Tam\"}",
				"jsonata": "\/* a comment and also a weird field *\/ $.{ \"niceField\": $.\"number #\" }",
				"output": ""
		}`,
			expectedOutput: "{\"Input\":\"{\\\"name\\\": \\\"Tam\\\"}\",\"Jsonata\":\"/* a comment and also a weird field */ $.{ \\\"niceField\\\": $.\\\"number #\\\" }\",\"Output\":\"{}\"}\n",
		},
		{
			method: http.MethodPost,
			name:   `blank`,
			testInput: `{
				"input": "{}",
				"jsonata": "$$",
				"output": ""
		}`,
			expectedOutput: "{\"Input\":\"{}\",\"Jsonata\":\"$$\",\"Output\":\"{}\"}\n",
		},
		{
			method: http.MethodPost,
			name:   `invalid input`,
			testInput: `{
				"input": "{",
				"jsonata": "$$",
				"output": ""
		}`,
			expectedOutput: "{\"Input\":\"{\",\"Jsonata\":\"$$\",\"Output\":\"input json error: unexpected end of JSON input\"}\n",
		},
		{
			method: http.MethodPost,
			name:   `invalid jsonata`,
			testInput: `{
				"input": "{}",
				"jsonata": "[$",
				"output": ""
		}`,
			expectedOutput: "{\"Input\":\"{}\",\"Jsonata\":\"[$\",\"Output\":\"jsonata error: could not compile [$: expected token ']' before end of expression\"}\n",
		},
		{
			method:         http.MethodGet,
			name:           `invalid method`,
			testInput:      `{}`,
			expectedOutput: "Request type other than POST not supported",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/whatever",
				ioutil.NopCloser(strings.NewReader(tt.testInput)))

			w := httptest.NewRecorder()

			jsonataRequest(w, req)

			res := w.Result()

			defer res.Body.Close()

			data, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}

			assert.Equal(t, string(data), tt.expectedOutput)
		})
	}
}
