package main_test

import (
	"backend/internal/testutil"
	task104 "backend/task104"
	"testing"
)

func TestRun(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, task104.Run)
}
