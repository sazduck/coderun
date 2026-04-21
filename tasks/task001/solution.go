package main

import (
	"bufio"
	"io"
	"os"
	"slices"
	"strconv"
)

func main() {
	Run(os.Stdin, os.Stdout)
}

func Run(r io.Reader, w io.Writer) error {
	s := bufio.NewScanner(r)
	buf := make([]byte, 0, 1<<16)
	s.Buffer(buf, 1<<20)
	s.Split(bufio.ScanWords)

	b := bufio.NewWriterSize(w, 1<<20)
	defer b.Flush()

	next := func() int {
		s.Scan()
		val, _ := strconv.Atoi(s.Text())
		return val
	}

	arr := make([]int, 3)
	for i := range 3 {
		arr[i] = next()
	}

	slices.Sort(arr)

	b.WriteString(strconv.Itoa(arr[1]))
	b.WriteByte('\n')
	return nil
}
