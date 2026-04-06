package main_test

import (
	"backend/internal/testutil"
	task101 "backend/task101"
	"testing"
)

func TestRun(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, task101.Run)
}
