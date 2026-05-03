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
	bw := bufio.NewWriter(w)
	defer bw.Flush()

	next := func() int {
		s.Scan()
		n, _ := strconv.Atoi(s.Text())
		return n
	}

	n := next()

	graph := make([][]int, n)

	for i := range graph {
		graph[i] = make([]int, n)
		for j := range graph[i] {
			graph[i][j] = next()
		}
	}

	start, target := next(), next()

	start--
	target--

	dist := make([]int, n)
	for i := range dist {
		dist[i] = -1
	}
	dist[target] = -1

	queue := []int{start}
	dist[start] = 0

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		for i := range n {
			if graph[v][i] == 1 && dist[i] == -1 {
				dist[i] = dist[v] + 1
				queue = append(queue, i)
			}
		}
	}
	bw.WriteString(strconv.Itoa(dist[target]) + "\n")

	return nil
}
