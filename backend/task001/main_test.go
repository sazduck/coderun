package task001_test

import (
	"backend/internal/testutil"
	"backend/task001"

	"testing"
)

func TestRun(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, task001.Run)
}
