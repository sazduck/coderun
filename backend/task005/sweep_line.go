package main

import (
	"bufio"
	"io"
	"os"
	"slices"
	"strconv"
)

func main() {
	Run(os.Stdin, os.Stdout)
}

func Run(r io.Reader, w io.Writer) error {
	segs, points := readInput(r)
	res := countPointsInSegs(points, segs)
	printRes(w, res)
	return nil
}

func printRes(w io.Writer, res []int) {
	buf := bufio.NewWriterSize(w, 1<<20)
	defer buf.Flush()

	b := make([]byte, 0, 20) // int64
	for _, v := range res {
		b = strconv.AppendInt(b[:0], int64(v), 10)
		buf.Write(b)
		buf.WriteByte(' ')
	}
	buf.WriteByte('\n')
}

type Segment struct {
	S int
	E int
}

const (
	begKind = -1
	endKind = 1
	pntKind = 0
)

type Event struct {
	pos  int
	kind int8
	idx  int
}

func countPointsInSegs(points []int, segs []Segment) []int {

	events := make([]Event, 0, 2*len(segs)+len(points))

	for _, s := range segs {
		events = append(events, Event{pos: s.S, kind: begKind})
		events = append(events, Event{pos: s.E, kind: endKind})
	}

	for i, p := range points {
		events = append(events, Event{pos: p, kind: pntKind, idx: i})
	}

	slices.SortFunc(events, func(a, b Event) int {
		if a.pos != b.pos {
			// return a.pos - b.pos
			if a.pos < b.pos {
				return -1
			}
			return 1
		}
		return int(a.kind - b.kind)
	})

	res := make([]int, len(points))

	// for i, point := range points {
	// 	for _, seg := range segs {
	// 		if seg.S <= point && point <= seg.E {
	// 			res[i] += 1
	// 		}
	// 	}
	// }

	segCount := 0

	for _, e := range events {
		switch e.kind {
		case begKind:
			segCount++
		case endKind:
			segCount--
		default:
			res[e.idx] = segCount
		}
	}

	return res
}

func readInput(r io.Reader) ([]Segment, []int) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)

	next := func() int {
		s.Scan()
		val, _ := strconv.Atoi(s.Text())
		return val
	}

	n, m := next(), next()

	segs := make([]Segment, n)
	for i := range n {
		a, b := next(), next()
		segs[i] = Segment{S: min(a, b), E: max(a, b)}
	}
	points := make([]int, m)
	for i := range m {
		points[i] = next()
	}

	return segs, points
}
