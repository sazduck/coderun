package main_test

import (
	"backend/internal/testutil"
	task008 "backend/task008"
	"testing"
)

func TestRun(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, task008.Run)
}
