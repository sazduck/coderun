package main_test

import (
	"testing"

	"github.com/sazduck/algo-practice-util/testutil"

	task002 "github.com/sazduck/coderun/tasks/backend/task002"
)

func TestRun(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, task002.Run)
}
