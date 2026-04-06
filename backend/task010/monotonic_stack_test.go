package main_test

import (
	"backend/internal/testutil"
	task010 "backend/task010"
	"testing"
)

func TestRun(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, task010.Run)
}
