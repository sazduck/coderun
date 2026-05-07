package main_test

import (
	"bytes"
	"strings"
	"testing"

	task009 "github.com/sazduck/coderun/tasks/backend/task009"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "example 0",
			input:  "2\r\nabcaz\r\n",
			output: "4",
		},
		{
			name:   "example 1",
			input:  "2\r\nhelto\r\n",
			output: "3",
		},
	}
	for _, tt := range tests {
		r := strings.NewReader(tt.input)
		var w bytes.Buffer

		t.Run(tt.name, func(t *testing.T) {
			task009.Run(r, &w)

			got := strings.TrimSpace(w.String())
			want := strings.TrimSpace(tt.output)

			if got != want {
				t.Error()
			}
		})
	}
}
