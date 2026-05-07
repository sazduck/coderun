package main_test

import (
	"bytes"
	"strings"
	"testing"

	task006 "github.com/sazduck/coderun/tasks/task006"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		// TODO
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			var w bytes.Buffer

			task006.Run(r, &w)

			got := strings.TrimSpace(w.String())
			want := strings.TrimSpace(tt.output)

			if got != want {
				t.Errorf("\ngot: %q\nwant: %q", got, want)
			}
		})
	}
}
