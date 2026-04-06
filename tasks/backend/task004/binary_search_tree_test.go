package main_test

import (
	"testing"

	task004 "github.com/sazduck/coderun/tasks/backend/task004"
	"github.com/sazduck/algo-practice-util/testutil"
)

func TestRun(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, task004.Run)
}
