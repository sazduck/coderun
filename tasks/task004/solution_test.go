package main_test

import (
	"testing"

	"github.com/sazduck/algo-practice-util/testutil"
	task004 "github.com/sazduck/coderun/tasks/task004"
)

func TestRun(t *testing.T) {
    testutil.RunWithDefaultTestCasesPath(t, task004.Run)
}
