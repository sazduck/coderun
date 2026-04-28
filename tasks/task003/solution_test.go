package main_test

import (
	"testing"

	"github.com/sazduck/algo-practice-util/testutil"
	task003 "github.com/sazduck/coderun/tasks/task003"
)

func TestRun(t *testing.T) {
    testutil.RunWithDefaultTestCasesPath(t, task003.Run)
}
