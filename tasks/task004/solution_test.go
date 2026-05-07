package main_test

import (
	"bytes"
	"strings"
	"testing"

	task004 "github.com/sazduck/coderun/tasks/task004"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "example 0",
			input:  "3 2\n",
			output: "1\n",
		},
		{
			name:   "example 1",
			input:  "31 34\n",
			output: "293930\n",
		},
	}
	for _, tt := range tests {
		r := strings.NewReader(tt.input)
		var w bytes.Buffer

		t.Run(tt.name, func(t *testing.T) {
			task004.Run(r, &w)

			got := strings.TrimSpace(w.String())
			want := strings.TrimSpace(tt.output)

			if got != want {
				t.Error()
			}
		})
	}
}
