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
	s.Buffer(make([]byte, 1<<20), 1<<20)
	s.Split(bufio.ScanWords)
	bw := bufio.NewWriterSize(w, 1<<20)
	defer bw.Flush()

	next := func() int {
		s.Scan()
		n, _ := strconv.Atoi(s.Text())
		return n
	}

	n := next()
	an := make([]int, n)
	for i := range an {
		an[i] = next()
	}

	m := next()
	am := make([]int, m)
	for i := range am {
		am[i] = next()
	}

	data := make([]int, (n+1)*(m+1))
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = data[i*(m+1) : (i+1)*(m+1)]
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if an[i-1] == am[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	lcs := make([]int, 0, max(n, m))
	i, j := n, m
	for i > 0 && j > 0 {
		if an[i-1] == am[j-1] {
			lcs = append(lcs, an[i-1])
			i--
			j--
		} else if dp[i-1][j] >= dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	for i := len(lcs) - 1; i >= 0; i-- {
		if i < len(lcs)-1 {
			bw.WriteString(" ")
		}
		bw.WriteString(strconv.Itoa(lcs[i]))
	}
	bw.WriteString("\n")
	return nil
}
