package main_test

import (
	"testing"

	task009 "github.com/sazduck/coderun/tasks/backend/task009"
	"github.com/sazduck/algo-practice-util/testutil"
)

func TestRun(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, task009.Run)
}
