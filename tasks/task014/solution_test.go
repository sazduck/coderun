package main_test

import (
	"bytes"
	"strings"
	"testing"

	task014 "github.com/sazduck/coderun/tasks/task014"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "Example 1",
			input:  "2 2 1 1 1\n2 2",
			output: "-1",
		},
		{
			name:   "Example 2",
			input:  "4 4 1 1 16\n1 1\n1 2\n1 3\n1 4\n2 1\n2 2\n2 3\n2 4\n3 1\n3 2\n3 3\n3 4\n4 1\n4 2\n4 3\n4 4",
			output: "42",
		},
	}
	for _, tt := range tests {
		r := strings.NewReader(tt.input)
		var w bytes.Buffer

		t.Run(tt.name, func(t *testing.T) {
			task014.Run(r, &w)

			got := strings.TrimSpace(w.String())
			want := strings.TrimSpace(tt.output)

			if got != want {
				t.Error()
			}
		})
	}
}
