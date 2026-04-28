package main

import (
	"bufio"
	"io"
	"os"
	"slices"
	"strconv"
)

func main() {
	// r := strings.NewReader("5 5\n9 9 9 9 9\n3 0 0 0 0\n9 9 9 9 9\n6 6 6 6 8\n9 9 9 9 9")
	r := os.Stdin
	w := os.Stdout

	Run(r, w)

}

func Run(r io.Reader, w io.Writer) error {

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)
	bw := bufio.NewWriter(w)
	defer bw.Flush()

	next := func() int {

		s.Scan()
		val, _ := strconv.Atoi(s.Text())
		return val
	}

	n := next()
	m := next()

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, m)
		for j := range matrix[i] {
			matrix[i][j] = next()
		}
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	dp[0][0] = matrix[0][0] // start

	for j := 1; j < m; j++ {
		dp[0][j] = dp[0][j-1] + matrix[0][j]
	}

	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + matrix[i][0]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = matrix[i][j] + max(dp[i-1][j], dp[i][j-1])
		}
	}

	// for _, v := range matrix {
	// 	fmt.Fprintf(bw, "%v\n", v)
	// }
	// bw.WriteByte('\n')
	// for _, v := range dp {
	// 	fmt.Fprintf(bw, "%v\n", v)
	// }

	path := make([]byte, 0, n+m-2)

	i, j := n-1, m-1

	for i > 0 || j > 0 {
		if i == 0 {
			path = append(path, 'R')
			j--
			continue
		}
		if j == 0 {
			path = append(path, 'D')
			i--
			continue
		}

		if dp[i-1][j] >= dp[i][j-1] {
			path = append(path, 'D')
			i--
			continue
		}

		path = append(path, 'R')
		j--
	}

	bw.WriteString(strconv.Itoa(dp[n-1][m-1]))
	bw.WriteByte('\n')

	for i := range len(path) / 2 {
		path[i], path[len(path)-1-i] = path[len(path)-1-i], path[i]
	}

	for _, v := range slices.Backward(path) {
		bw.WriteByte(v)
		bw.WriteByte(' ')
	}
	bw.WriteByte('\n')

	return nil
}
