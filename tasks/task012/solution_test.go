package main_test

import (
	"testing"

	"github.com/sazduck/algo-practice-util/testutil"
	task012 "github.com/sazduck/coderun/tasks/task012"
)

func TestRun(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, task012.Run)
}
