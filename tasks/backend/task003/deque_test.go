package main_test

import (
	"testing"

	task003 "github.com/sazduck/coderun/tasks/backend/task003"
	"github.com/sazduck/algo-practice-util/testutil"
)

func TestRun(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, task003.Run)
}
