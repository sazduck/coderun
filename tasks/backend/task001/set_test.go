package main_test

import (
	"bytes"
	"strings"
	"testing"

	task001 "github.com/sazduck/coderun/tasks/backend/task001"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "example 0",
			input:  "She sells sea shells on the sea shore;\nThe shells that she sells are sea shells I'm sure.\nSo if she sells sea shells on the sea shore,\nI'm sure that the shells are sea shore shells.\n",
			output: "19\n",
		},
		{
			name:   "example 1",
			input:  "AA aa Aa aA",
			output: "4\n",
		},
		{
			name:   "example 2",
			input:  "a a a a a a a a a a a a a a a a a a a a a a a a a a a a a a a a a a a a a\n",
			output: "1\n",
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
				t.Errorf("\ngot:  %q\nwant: %q", got, want)
			}
		})
	}
}
