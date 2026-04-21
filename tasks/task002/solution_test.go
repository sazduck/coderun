package main_test

import (
	"testing"

	"github.com/sazduck/algo-practice-util/testutil"
	task002 "github.com/sazduck/coderun/tasks/task002"
)

func TestRun(t *testing.T) {
	testutil.Run(t, []testutil.TestCase{
		{
			Name:     "example 0",
			Input:    "5 5\n1 1 1 1 1\n3 100 100 100 100\n1 1 1 1 1\n2 2 2 2 1\n1 1 1 1 1\n",
			Expected: "11\n",
		},
	}, task002.Run)
}
