package main_test

import (
	"testing"

	task006 "github.com/sazduck/coderun/tasks/backend/task006"
	"github.com/sazduck/algo-practice-util/testutil"
)

func TestRun(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, task006.Run)
}
