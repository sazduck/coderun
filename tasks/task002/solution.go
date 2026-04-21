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
	b := bufio.NewWriterSize(w, 1<<20)
	defer b.Flush()

	next := func() int {
		s.Scan()
		val, _ := strconv.Atoi(s.Text())
		return val
	}

	n, m := next(), next()

	arr := make([][]int, n)
	for i := range n {
		arr[i] = make([]int, m)
		for j := range m {
			cost := next()
			curEl := &arr[i][j]

			if i == 0 && j == 0 {
				*curEl = cost
			} else if i == 0 {
				*curEl = cost + arr[i][j-1]
			} else if j == 0 {
				*curEl = cost + arr[i-1][j]
			} else {
				*curEl = cost + min(arr[i-1][j], arr[i][j-1])
			}
		}
	}

	b.WriteString(strconv.Itoa(arr[n-1][m-1]))
	b.WriteByte('\n')
	return nil
}
