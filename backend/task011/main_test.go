package main_test

import (
	"backend/internal/testutil"
	task011 "backend/task011"
	"testing"
)

func TestRun(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, task011.Run)
}
