package main_test

import (
	"testing"

	"github.com/sazduck/algo-practice-util/testutil"

	task001 "github.com/sazduck/coderun/tasks/backend/task001"
)

func TestRun(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, task001.Run)
}
