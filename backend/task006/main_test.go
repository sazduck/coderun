package main_test

import (
	"backend/internal/testutil"
	task006 "backend/task006"
	"testing"
)

func TestRun(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, task006.Run)
}
