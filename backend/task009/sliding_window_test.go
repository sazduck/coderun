package main_test

import (
	"backend/internal/testutil"
	task009 "backend/task009"
	"testing"
)

func TestRun(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, task009.Run)
}
