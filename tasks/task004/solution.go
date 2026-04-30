package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {

	r := os.Stdin
	w := os.Stdout

	Run(r, w)
}

func Run(r io.Reader, w io.Writer) error {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)

	bw := bufio.NewWriterSize(w, 1<<20)

	defer bw.Flush()

	next := func() int {
		s.Scan()
		n, _ := strconv.Atoi(s.Text())
		return n
	}

	n := next()
	m := next()

	/*
		d 2
		r 1
		d 1
		r 2
	*/

	board := make([][]int, n+1)
	for i := range board {
		board[i] = make([]int, m+1)
	}

	board[1][1] = 1

	for i := 2; i < n+1; i++ {

		for j := 2; j < m+1; j++ {

			if j >= 3 {
				board[i][j] += board[i-1][j-2]
			}

			if i >= 3 {
				board[i][j] += board[i-2][j-1]
			}

		}
	}

	bw.WriteString(strconv.Itoa(board[n][m]))
	bw.WriteByte('\n')

	return nil
}
