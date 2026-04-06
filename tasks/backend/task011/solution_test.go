package main_test

import (
	"testing"

	task011 "github.com/sazduck/coderun/tasks/backend/task011"
	"github.com/sazduck/algo-practice-util/testutil"
)

func TestRun(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, task011.Run)
}
