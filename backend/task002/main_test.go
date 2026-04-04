package task002_test

import (
	"backend/internal/testutil"
	"backend/task002"

	"testing"
)

func TestRun(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, task002.Run)
}
