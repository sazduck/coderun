package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	w := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer w.Flush()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords) // Магия: сам убирает пробелы, \n, \t и \r

	wordSet := make(map[string]struct{})

	for scanner.Scan() {
		wordSet[scanner.Text()] = struct{}{}
	}

	wordCount := len(wordSet)

	w.WriteString(strconv.Itoa(wordCount))
	w.WriteByte('\n')
}
