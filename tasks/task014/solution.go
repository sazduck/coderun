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

type coord struct{ r, c int }

func (p coord) isValid(n, m int) bool {
	return p.r >= 0 && p.r < n && p.c >= 0 && p.c < m
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

	n, m := next(), next()

	feeder := coord{next() - 1, next() - 1}
	q := next()

	fleas := make([]coord, q)
	for i := range fleas {
		fleas[i] = coord{next() - 1, next() - 1}
	}

	grid := make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, m)
		for j := range grid[i] {
			grid[i][j] = -1
		}
	}

	moves := [...]coord{
		{-2, -1}, {-2, +1},
		{-1, -2}, {-1, +2},
		{+1, -2}, {+1, +2},
		{+2, -1}, {+2, +1},
	}

	queue := make([]coord, 0, n*m)
	grid[feeder.r][feeder.c] = 0
	queue = append(queue, feeder)

	head := 0
	for head < len(queue) {
		curr := queue[head]
		head++

		for _, move := range moves {

			next := coord{curr.r + move.r, curr.c + move.c}
			if !next.isValid(n, m) || grid[next.r][next.c] != -1 {
				continue
			}
			grid[next.r][next.c] = grid[curr.r][curr.c] + 1
			queue = append(queue, next)
		}
	}

	totalJumps := 0
	for _, f := range fleas {
		if grid[f.r][f.c] == -1 {
			bw.WriteString(strconv.Itoa(-1))
			bw.WriteByte('\n')
			return nil
		}
		totalJumps += grid[f.r][f.c]
	}
	bw.WriteString(strconv.Itoa(totalJumps))
	bw.WriteByte('\n')
	return nil
}
