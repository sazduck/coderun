package main_test

import (
	"bytes"
	"strings"
	"testing"

	task003 "github.com/sazduck/coderun/tasks/task003"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "example 0",
			input:  "5 5\n9 9 9 9 9\n3 0 0 0 0\n9 9 9 9 9\n6 6 6 6 8\n9 9 9 9 9\n",
			output: "74\nD D R R R R D D\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			var w bytes.Buffer

			task003.Run(r, &w)

			got := strings.TrimSpace(w.String())
			want := strings.TrimSpace(tt.output)
			if got != want {
				t.Errorf("\ngot: %q\nwant: %q", got, want)
			}
		})
	}
}
