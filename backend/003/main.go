package main

import (
	"bufio"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	res, err := FindMinimum(os.Stdin)
	if err != nil {
		return
	}
	strNums := make([]string, len(res))
	for i, v := range res {
		strNums[i] = strconv.Itoa(v)
	}
	ans := strings.Join(strNums, "\n")

	w := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer w.Flush()
	w.WriteString(ans)
	w.WriteByte('\n')
}

func FindMinimum(r io.Reader) ([]int, error) {
	s := bufio.NewScanner(r)

	s.Scan()
	nums := strings.Fields(s.Text())

	seqSize := 0
	winSize := 0

	if len(nums) >= 2 {
		if n, err := strconv.Atoi(nums[0]); err == nil {
			seqSize = n
		} else {
			return nil, err
		}
		if k, err := strconv.Atoi(nums[1]); err == nil {
			winSize = k
		} else {
			return nil, err
		}
	}

	seq := make([]int, 0, seqSize)
	s.Scan()
	for s := range strings.FieldsSeq(s.Text()) {
		if n, err := strconv.Atoi(s); err == nil {
			seq = append(seq, n)
		} else {
			return nil, err
		}

	}

	result := make([]int, 0, seqSize-winSize)
	for i := 0; i <= seqSize-winSize; i++ {
		result = append(result, slices.Min(seq[i:i+winSize]))
	}

	return result, nil

}
