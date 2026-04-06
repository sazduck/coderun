package main_test

import (
	"backend/internal/testutil"
	task103 "backend/task103"
	"testing"
)

func TestRun(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, task103.Run)
}
