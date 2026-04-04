package task001

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	Run(os.Stdin, os.Stdout)
}

func Run(r io.Reader, w io.Writer) error {
	count := CountUniqueWords(r)

	buf := bufio.NewWriterSize(w, 1<<20)
	defer buf.Flush()
	buf.WriteString(strconv.Itoa(count))
	buf.WriteByte('\n')
	return nil
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
