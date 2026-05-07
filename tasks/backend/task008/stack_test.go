package main_test

import (
	"bytes"
	"strings"
	"testing"

	task008 "github.com/sazduck/coderun/tasks/backend/task008"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "example 0",
			input:  "()[]\n",
			output: "yes\n",
		},
		{
			name:   "example 1",
			input:  "([)]\n",
			output: "no\n",
		},
		{
			name:   "example 2",
			input:  "(\n",
			output: "no\n",
		},
	}
	for _, tt := range tests {
		r := strings.NewReader(tt.input)
		var w bytes.Buffer

		t.Run(tt.name, func(t *testing.T) {
			task008.Run(r, &w)

			got := strings.TrimSpace(w.String())
			want := strings.TrimSpace(tt.output)

			if got != want {
				t.Error()
			}
		})
	}
}
