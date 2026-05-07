package main_test

import (
	"bytes"
	"strings"
	"testing"

	task002 "github.com/sazduck/coderun/tasks/backend/task002"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "example 0",
			input:  "3\nHello Hi\nBye Goodbye\nList Array\nGoodbye\n",
			output: "Bye\n",
		},
		{
			name:   "example 1",
			input:  "1\nbeep Car\nCar\n",
			output: "beep\n",
		},
		{
			name:   "example 2",
			input:  "2\nOlolo Ololo\nNumbers 1234567890\nNumbers\n",
			output: "1234567890\n",
		},
	}
	for _, tt := range tests {
		r := strings.NewReader(tt.input)
		var w bytes.Buffer

		t.Run(tt.name, func(t *testing.T) {
			task002.Run(r, &w)

			got := strings.TrimSpace(w.String())
			want := strings.TrimSpace(tt.output)

			if got != want {
				t.Error()
			}
		})
	}
}
