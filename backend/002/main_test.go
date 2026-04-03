package main_test

import (
	main "backend/002"
	"backend/internal/testutil"
	"io"
	"testing"
)

func TestCountUniqueWords(t *testing.T) {
	testutil.RunWithDefaultTestCasesPath(t, func(r io.Reader) string {
		str, err := main.PickSynonym(r)
		if err != nil {
			t.Fatalf("ошибка при выборе синонима")
		}
		return str
	})
}
