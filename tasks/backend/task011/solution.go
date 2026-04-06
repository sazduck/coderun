package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
)

func main() {
	Run(os.Stdin, os.Stdout)
}

func Run(r io.Reader, w io.Writer) error {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)

	ns := make([]int, 0, 1024)

	for s.Scan() {
		val, _ := strconv.Atoi(s.Text())
		ns = append(ns, val)
	}

	n1, n2 := countRes(ns)

	fmt.Fprintf(w, "%d %d", n1, n2)
	return nil
}

func countRes(nums []int) (int, int) {
	slices.Sort(nums)
	n := len(nums)

	posProd := nums[n-1] * nums[n-2]
	negProd := nums[0] * nums[1]

	if posProd > negProd {
		return nums[n-2], nums[n-1]

	}
	return nums[0], nums[1]
}
