package main_test

import (
	"bytes"
	"strings"
	"testing"

	task005 "github.com/sazduck/coderun/tasks/backend/task005"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "example 0",
			input:  "3 2\n0 5\n-3 2\n7 10\n1 6\n",
			output: "2 0 \n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			var w bytes.Buffer

			task005.Run(r, &w)

			got := strings.TrimSpace(w.String())
			want := strings.TrimSpace(tt.output)

			if got != want {
				t.Errorf("\ngot:  %q\nwant: %q", got, want)
			}
		})
	}
}
