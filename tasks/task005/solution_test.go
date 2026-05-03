package main_test

import (
	"testing"

	"github.com/sazduck/algo-practice-util/testutil"
	task005 "github.com/sazduck/coderun/tasks/task005"
)

func TestRun(t *testing.T) {
    testutil.RunWithDefaultTestCasesPath(t, task005.Run)
}
