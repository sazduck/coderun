package main_test

import (
	"backend/internal/testutil"
	task007 "backend/task007"
	"testing"
)

func TestRun(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, task007.Run)
}
