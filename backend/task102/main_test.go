package main_test

import (
	"backend/internal/testutil"
	task102 "backend/task102"
	"testing"
)

func TestRun(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, task102.Run)
}
