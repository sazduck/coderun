package main

import (
	"bufio"
	"cmp"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

func main() {
	Run(os.Stdin, os.Stdout)
}

type Interval struct {
	L, R, Len int
}

func Run(r io.Reader, w io.Writer) error {
	s, err := readInput(r)
	if err != nil {
		return err
	}

	intervals := findIntervals(s)
	res, totalLen := solveIntervals(intervals)

	b := bufio.NewWriterSize(w, 1<<20)
	defer b.Flush()
	fmt.Fprintf(b, "%d %d\n", len(res), totalLen)
	for _, iv := range res {
		fmt.Fprintf(b, "%d %d\n", iv.L, iv.R)
	}
	return nil
}

func readInput(r io.Reader) (string, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

func findIntervals(s string) []Interval {
	n := len(s)
	intervals := make([]Interval, 0, n/4)

	for i := range n {
		c := s[i]
		if c != '{' && c != '[' && c != '"' && !(c >= '0' && c <= '9') && c != '-' {
			continue
		}

		jd := json.NewDecoder(strings.NewReader(s[i:]))

		var v json.RawMessage
		if jd.Decode(&v) == nil {
			end := i + int(jd.InputOffset())
			intervals = append(intervals, Interval{i, end, end - i})
		}
	}
	return intervals
}

func solveIntervals(intervals []Interval) ([]Interval, int) {
	m := len(intervals)
	if m == 0 {
		return nil, 0
	}

	slices.SortFunc(intervals, func(a, b Interval) int {
		if c := cmp.Compare(a.R, b.R); c != 0 {
			return c
		}
		return cmp.Compare(a.L, b.L)
	})

	ends := make([]int, m)
	for i, iv := range intervals {
		ends[i] = iv.R
	}

	dp := make([]int, m+1)
	prev := make([]int, m+1)

	for i := range m {
		iv := intervals[i]

		idx, _ := slices.BinarySearch(ends[:i], iv.L)
		for idx < i && ends[idx] == iv.L {
			idx++
		}

		take := iv.Len + dp[idx]
		if take > dp[i] {
			dp[i+1] = take
			prev[i+1] = idx
		} else {
			dp[i+1] = dp[i]
			prev[i+1] = i
		}
	}

	res := make([]Interval, 0, m/2)
	curr := m
	for curr > 0 {
		if dp[curr] == dp[curr-1] {
			curr--
		} else {
			res = append(res, intervals[curr-1])
			curr = prev[curr]
		}
	}

	return res, dp[m]
}

/*
output:

hello "a"bbbbbbbbb"c" world
        ^8         ^19       8-19

1      11
^ocur  ^len
8      19
^start ^end

data: {"key": "value"} and [1, 2, 3]
      ^6              ^22  ^27      ^36

no json {} {"x": 1} ??? {"y": [2, 3]} end
        ^ ^^11     ^19  ^24          ^37

*/
