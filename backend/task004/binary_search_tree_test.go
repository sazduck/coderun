package main_test

import (
	"backend/internal/testutil"
	task004 "backend/task004"
	"testing"
)

func TestRun(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, task004.Run)
}
