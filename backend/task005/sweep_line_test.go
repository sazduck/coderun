package main_test

import (
	"backend/internal/testutil"
	task005 "backend/task005"
	"testing"
)

func TestRun(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, task005.Run)
}
