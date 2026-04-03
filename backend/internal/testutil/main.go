package testutil

import (
	"encoding/json"
	"io"
	"os"
	"strings"
	"testing"
)

// JSON.stringify(
//   [...document.querySelectorAll('.io-sample>div')].map((d, i) => ({
//     name: `example ${i}`,
//     input: d.querySelector('.input-snippet code').innerText.trim(),
//     expected: d.querySelector('.output-snippet code').innerText.trim(),
//   }))
// );

type TestCase struct {
	Name     string `json:"name"`
	Input    string `json:"input"`
	Expected string `json:"expected"`
}

func RunTests(t *testing.T, tc []TestCase, f func(io.Reader) string) {
	for _, tt := range tc {
		t.Run(tt.Name, func(t *testing.T) {
			r := strings.NewReader(tt.Input)
			res := f(r)
			if res != tt.Expected {
				t.Errorf("Ввод: %s, ожидалось: %s, получено: %s", tt.Input, tt.Expected, res)
			}
		})
	}
}

func LoadTestCases(t *testing.T, path string) []TestCase {
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Не удалось прочитать файл %s: %v", path, err)
	}
	var tests []TestCase
	err = json.Unmarshal(data, &tests)
	if err != nil {
		t.Fatalf("Ошибка при десериализации JSON %v", err)
	}
	return tests
}

func RunWithDefaultTestCasesPath(t *testing.T, f func(io.Reader) string) {
	const defaultPath = "./test_cases.json"
	tests := LoadTestCases(t, defaultPath)
	RunTests(t, tests, f)
}
