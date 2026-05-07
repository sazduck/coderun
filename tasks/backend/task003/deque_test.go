package main_test

import (
	"bytes"
	"strings"
	"testing"

	task003 "github.com/sazduck/coderun/tasks/backend/task003"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "example 0",
			input:  "7 3\n1 3 2 4 5 3 1\n",
			output: "1\n2\n2\n3\n1\n",
		},
	}
	for _, tt := range tests {
		r := strings.NewReader(tt.input)
		var w bytes.Buffer

		t.Run(tt.name, func(t *testing.T) {
			task003.Run(r, &w)

			got := strings.TrimSpace(w.String())
			want := strings.TrimSpace(tt.output)

			if got != want {
				t.Error()
			}
		})
	}
}
