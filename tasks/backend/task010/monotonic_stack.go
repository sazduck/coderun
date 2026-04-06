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
	heights := readInput(r)
	res := largestRectangleArea(heights)
	return writeRes(w, res)
}

func writeRes(w io.Writer, n int) error {
	b := bufio.NewWriterSize(w, 1<<20)
	defer b.Flush()
	b.WriteString(strconv.Itoa(n))
	b.WriteByte('\n')
	return nil
}

func readInput(r io.Reader) []int {
	s := bufio.NewScanner(r)
	buf := make([]byte, 1<<20)
	s.Buffer(buf, 1<<20)
	s.Split(bufio.ScanWords)

	next := func() int {
		s.Scan()
		val, _ := strconv.Atoi(string(s.Bytes()))
		return val
	}

	n := next()

	heights := make([]int, n)
	for i := range n {
		heights[i] = next()
	}
	return heights
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	stack := make([]int, n+1)
	top := -1
	maxArea := 0

	for i := range n {
		currentHeight := 0
		if i < n {
			currentHeight = heights[i]
		}
		for top >= 0 && currentHeight < heights[stack[top]] {
			h := heights[stack[top]]
			top-- // Pop

			w := i
			if top >= 0 {
				w = i - stack[top] - 1
			}

			area := h * w
			if area > maxArea {
				maxArea = area
			}
		}
		top++
		stack[top] = i
	}

	return maxArea
}
