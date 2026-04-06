package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

	// r := strings.NewReader("1\n")
	// Run(r, os.Stdout)
	Run(os.Stdin, os.Stdout)
}

func Run(r io.Reader, w io.Writer) error {
	line, _ := bufio.NewReader(r).ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line)) // n <= 35

	c := countSeqs(n)

	b := bufio.NewWriter(w)
	defer b.Flush()
	b.WriteString(strconv.Itoa(c))
	b.WriteByte('\n')
	return nil
}
func countSeqs(n int) int {
	switch n {
	case 1:
		return 2 // 0 | 1 => 2
	case 2:
		return 4 // 00 | 01 | 10 | 11 => 4
	case 3:
		return 7 // 000 | 001 | 010 | 011 | 100 | 101 | 110 | 111 => 7
	}

	a, b, c := 2, 4, 7
	for i := 4; i <= n; i++ {
		a, b, c = b, c, a+b+c // сдвигаем "окно" и считаем новое значение
	}
	return c
}
