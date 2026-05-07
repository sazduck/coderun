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
	res, err := FindMinimum(r)
	if err != nil {
		return err
	}

	buf := bufio.NewWriterSize(w, 1<<20)
	defer buf.Flush()
	for i := range res {
		buf.WriteString(strconv.Itoa(res[i]))
		buf.WriteByte('\n')
	}
	return nil
}

func FindMinimum(r io.Reader) ([]int, error) {

	winSize, seq := parseInput(r)
	result := minByWindow(seq, winSize)

	return result, nil

}

func parseInput(r io.Reader) (int, []int) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)

	next := func() int {
		s.Scan()
		val, _ := strconv.Atoi(s.Text())
		return val
	}

	seqSize := next()
	winSize := next()

	seq := make([]int, seqSize)
	for i := range seqSize {
		seq[i] = next()
	}
	return winSize, seq

}

func minByWindow(seq []int, winSize int) []int {

	if len(seq) == 0 || winSize <= 0 {
		return nil
	}

	resCount := len(seq) - winSize + 1
	res := make([]int, 0, resCount)
	deque := make([]int, 0, winSize) // Храним индексы

	for i, num := range seq {
		leftBoundary := i - winSize

		if len(deque) > 0 && deque[0] <= leftBoundary {
			deque = deque[1:]
		}

		for len(deque) > 0 {
			lastIndex := deque[len(deque)-1]
			if seq[lastIndex] < num {
				break
			}
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)

		if i >= winSize-1 {
			minIndex := deque[0]
			res = append(res, seq[minIndex])
		}
	}
	return res
}
