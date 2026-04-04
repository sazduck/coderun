package task003_test

import (
	"backend/internal/testutil"
	"backend/task003"
	"testing"
)

func TestRun(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, task003.Run)
}
