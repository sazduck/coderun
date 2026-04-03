package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	count := CountUniqueWords(os.Stdin)

	w := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer w.Flush()
	w.WriteString(strconv.Itoa(count))
	w.WriteByte('\n')
}

// CountUniqueWords принимает [io.Reader], который преобрузается в [bufio.Scanner],
// который читает количество уникльных слов разеделенных \n, \t, \s
// и возвращает их количество
func CountUniqueWords(r io.Reader) int {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)

	wordSet := make(map[string]struct{})
	for s.Scan() {
		wordSet[s.Text()] = struct{}{}
	}
	return len(wordSet)
}
