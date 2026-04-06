package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	// r := strings.NewReader("2\nasdf\n")
	// Run(r, os.Stdout)
	Run(os.Stdin, os.Stdout)
}

func Run(r io.Reader, w io.Writer) error {
	s := bufio.NewScanner(r)
	const maxBuf = 1 << 20
	s.Buffer(make([]byte, maxBuf), maxBuf)

	if !s.Scan() {
		return nil
	}
	k, _ := strconv.Atoi(s.Text())

	if !s.Scan() {
		return nil
	}
	line := s.Bytes()

	globalMax := 0

	for char := byte('a'); char <= 'z'; char++ {
		cnt := 0

		j := 0
		for i := range line {
			if line[i] != char {
				cnt++
			}

			for cnt > k {
				if line[j] != char {
					cnt--
				}
				j++
			}

			currentLen := i - j + 1
			if currentLen > globalMax {
				globalMax = currentLen
			}
		}
	}

	var resBuf [20]byte
	w.Write(strconv.AppendInt(resBuf[:0], int64(globalMax), 10))
	w.Write([]byte("\n"))

	return s.Err()
}
