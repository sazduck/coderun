package main_test

import (
	"bytes"
	"strings"
	"testing"

	task002 "github.com/sazduck/coderun/tasks/task002"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "example 0",
			input:  "5 5\n1 1 1 1 1\n3 100 100 100 100\n1 1 1 1 1\n2 2 2 2 1\n1 1 1 1 1\n",
			output: "11\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			var w bytes.Buffer

			task002.Run(r, &w)

			got := strings.TrimSpace(w.String())
			want := strings.TrimSpace(tt.output)
			if got != want {
				t.Errorf("\ngot: %q\nwant: %q", got, want)
			}
		})
	}
}
