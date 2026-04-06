package main_test

import (
	"testing"

	task008 "github.com/sazduck/coderun/tasks/backend/task008"
	"github.com/sazduck/algo-practice-util/testutil"
)

func TestRun(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, task008.Run)
}
