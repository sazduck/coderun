package main_test

import (
	"testing"

	task010 "github.com/sazduck/coderun/tasks/backend/task010"
	"github.com/sazduck/algo-practice-util/testutil"
)

func TestRun(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, task010.Run)
}
