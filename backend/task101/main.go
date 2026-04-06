package main

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
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)

	next := func() int {
		s.Scan()
		val, _ := strconv.Atoi(s.Text())
		return val
	}

	n, k := next(), next()

	count := 0
	for range n {
		val := next()
		if val < k {
			break
		}
		count++

	}

	b := bufio.NewWriterSize(w, 1<<20)
	defer b.Flush()
	b.WriteString(strconv.Itoa(count))
	b.WriteByte('\n')
	return nil
}
