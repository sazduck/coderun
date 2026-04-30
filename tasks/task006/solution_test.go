package main_test

import (
	"testing"

	"github.com/sazduck/algo-practice-util/testutil"
	task006 "github.com/sazduck/coderun/tasks/task006"
)

func TestRun(t *testing.T) {
    testutil.RunWithDefaultTestCasesPath(t, task006.Run)
}
