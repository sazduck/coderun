package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func main() {
	Run(os.Stdin, os.Stdout)
}

func Run(r io.Reader, w io.Writer) error {

	line, _ := bufio.NewReaderSize(r, 1<<20).ReadString('\n')
	line = strings.TrimSpace(line)

	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := make([]rune, 0, len(line))

	res := "yes"
	for _, char := range line {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
			continue
		}
		opening, isClosing := pairs[char]
		if !isClosing {
			continue
		}
		if len(stack) == 0 || stack[len(stack)-1] != opening {
			res = "no"
			stack = nil
			break
		}
		stack = stack[:len(stack)-1]

	}
	if len(stack) != 0 {
		res = "no"
	}

	b := bufio.NewWriter(w)
	defer b.Flush()
	b.WriteString(res)
	b.WriteByte('\n')
	return nil
}
