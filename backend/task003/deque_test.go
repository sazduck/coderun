package main_test

import (
	"backend/internal/testutil"
	task003 "backend/task003"
	"testing"
)

func TestRun(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, task003.Run)
}
