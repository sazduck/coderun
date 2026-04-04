package task003

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

	seqSize, winSize, seq, err := parseInput(r)
	if err != nil {
		return nil, err
	}
	result := minByWindow(seq, seqSize, winSize)

	return result, nil

}

func parseInput(r io.Reader) (int, int, []int, error) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)
	s.Scan()
	seqSize, err := strconv.Atoi(s.Text())
	if err != nil {
		return 0, 0, nil, err
	}

	s.Scan()
	winSize, err := strconv.Atoi(s.Text())
	if err != nil {
		return 0, 0, nil, err
	}

	seq := make([]int, 0, seqSize)
	for s.Scan() {
		if n, err := strconv.Atoi(s.Text()); err == nil {
			seq = append(seq, n)
		} else {
			return 0, 0, nil, err
		}

	}
	return seqSize, winSize, seq, err

}

func minByWindow(seq []int, seqSize, winSize int) []int {
	if seqSize == 0 || winSize <= 0 {
		return nil
	}

	resCount := seqSize - winSize + 1
	res := make([]int, 0, resCount)
	deque := make([]int, 0, winSize) // Храним индексы

	for i := 0; i < seqSize; i++ {
		leftBoundary := i - winSize

		if len(deque) > 0 && deque[0] <= leftBoundary {
			deque = deque[1:]
		}

		for len(deque) > 0 {
			lastIndex := deque[len(deque)-1]
			if seq[lastIndex] < seq[i] {
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
