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
	next := func() int {
		s.Scan()
		val, _ := strconv.Atoi(s.Text())
		return val
	}

	n := next()
	aCount := make([]int, n)
	for i := range n {
		aCount[i] = next()
	}

	var totalGoodness int64 = 0
	for i := range n - 1 {
		totalGoodness += int64(min(aCount[i], aCount[i+1]))
	}

	buf := bufio.NewWriterSize(w, 1<<20)
	defer buf.Flush()
	_, err := buf.WriteString(strconv.FormatInt(totalGoodness, 10) + "\n")

	return err
}
