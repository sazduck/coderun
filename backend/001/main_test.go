package main_test

import (
	main "backend/001"
	"backend/internal/testutil"
	"io"
	"strconv"
	"testing"
)

func TestPickSynonym(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, func(r io.Reader) string {
		return strconv.Itoa(main.CountUniqueWords(r))
	})
}
