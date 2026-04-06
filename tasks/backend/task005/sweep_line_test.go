package main_test

import (
	"testing"

	task005 "github.com/sazduck/coderun/tasks/backend/task005"
	"github.com/sazduck/algo-practice-util/testutil"
)

func TestRun(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, task005.Run)
}
