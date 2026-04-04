package task004_test

import (
	"backend/internal/testutil"
	"backend/task004"
	"testing"
)

func TestRun(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, task004.Run)
}
