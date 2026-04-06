package main_test

import (
	"backend/internal/testutil"
	task002 "backend/task002"

	"testing"
)

func TestRun(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, task002.Run)
}
