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

const MaxN = 1_000_000

func Run(r io.Reader, w io.Writer) error {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)

	scan := func() int {
		s.Scan()
		val, _ := strconv.Atoi(s.Text())
		return val
	}

	n := scan()

	nums := make([]int, n)
	for i := range n {
		nums[i] = scan()
	}

	res := make([]int, 0, 2)
	for i := 0; i < n-1; i++ {
		curr, next := nums[i], nums[i+1]

		if curr > next {
			res = append(res, i+1)
			if curr-next > n-1 {
				res = append(res, i+1)
			}
		}
		if len(res) >= 2 {
			break
		}
	}

	b := bufio.NewWriterSize(w, 1<<20)
	defer b.Flush()
	b.WriteString(strconv.Itoa(res[0]) + " " + strconv.Itoa(res[1]) + "\n")
	return nil
}
