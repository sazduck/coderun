package main_test

import (
	"backend/internal/testutil"
	task001 "backend/task001"

	"testing"
)

func TestRun(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, task001.Run)
}
