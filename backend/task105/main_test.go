package main_test

import (
	"backend/internal/testutil"
	task105 "backend/task105"
	"testing"
)

func TestRun(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, task105.Run)
}
