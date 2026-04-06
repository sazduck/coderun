package main_test

import (
	"testing"

	task007 "github.com/sazduck/coderun/tasks/backend/task007"
	"github.com/sazduck/algo-practice-util/testutil"
)

func TestRun(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, task007.Run)
}
