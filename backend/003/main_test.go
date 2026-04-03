package main_test

import (
	main "backend/003"
	"backend/internal/testutil"
	"io"
	"strconv"
	"strings"
	"testing"
)

func TestFindMinimum(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, func(r io.Reader) string {

		res, err := main.FindMinimum(r)
		if err != nil {
			t.Fatalf("error: %v", err)
		}
		strNums := make([]string, len(res))
		for i, v := range res {
			strNums[i] = strconv.Itoa(v)
		}
		return strings.Join(strNums, "\n")
	})
}
