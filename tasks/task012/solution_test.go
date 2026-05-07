package main_test

import (
	"bytes"
	"strings"
	"testing"

	task012 "github.com/sazduck/coderun/tasks/task012"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "example 0",
			input:  "3\n1 2 3\n3\n2 3 1\n",
			output: "2 3 \n",
		},
		{
			name:   "example 1",
			input:  "3\n1 2 3\n3\n3 2 1\n",
			output: "1 \n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			var w bytes.Buffer

			task012.Run(r, &w)

			got := strings.TrimSpace(w.String())
			want := strings.TrimSpace(tt.output)

			if got != want {
				t.Errorf("\ngot: %q\nwant: %q", got, want)
			}
		})
	}
}
