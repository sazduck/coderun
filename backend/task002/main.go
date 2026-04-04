package task002

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

// Программа получает на вход количество пар синонимов N.
// Далее следует N строк, каждая строка содержит ровно два слова-синонима.
// После этого следует одно слово.
func main() {
	Run(os.Stdin, os.Stdout)
}

func Run(r io.Reader, w io.Writer) error {
	ans, err := PickSynonym(r)
	if err != nil {
		return err
	}

	buf := bufio.NewWriterSize(w, 1<<20)
	defer buf.Flush()
	buf.WriteString(ans)
	buf.WriteByte('\n')
	return nil
}

func PickSynonym(r io.Reader) (string, error) {
	s := bufio.NewScanner(r)

	if !s.Scan() {
		return "", errors.New("не найдено строк")
	}
	n, _ := strconv.Atoi(s.Text())

	synonyms := make(map[string]string, n*2)
	for range n {
		if !s.Scan() {
			break
		}
		w := strings.Fields(s.Text())
		if len(w) >= 2 {
			synonyms[w[0]] = w[1]
			synonyms[w[1]] = w[0]
		}
	}
	s.Scan()
	needle := strings.TrimSpace(s.Text())
	return synonyms[needle], nil

}
