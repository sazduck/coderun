package main_test

import (
	"bytes"
	"strings"
	"testing"

	task001 "github.com/sazduck/coderun/tasks/task001"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "Example 1",
			input:  "1 2 3",
			output: "2",
		},
		{
			name:   "Example 2",
			input:  "1000 -1000 0",
			output: "0",
		},
		{
			name:   "Example 3",
			input:  "3 1 3",
			output: "3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := strings.NewReader(tt.input)
			var w bytes.Buffer

			task001.Run(r, &w)

			got := strings.TrimSpace(w.String())
			want := strings.TrimSpace(tt.output)
			if got != want {
				t.Errorf("\ngot: %q\nwant: %q", got, want)
			}
		})
	}
}
