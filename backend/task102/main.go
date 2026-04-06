package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	Run(os.Stdin, os.Stdout)
}

const MaxN = 1_000_000

func Run(r io.Reader, w io.Writer) error {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)
	bw := bufio.NewWriterSize(w, 1<<20)
	defer bw.Flush()

	next := func() int {
		s.Scan()
		val, _ := strconv.Atoi(s.Text())
		return val
	}
	n := next()

	res := getMinPrimes(n)

	resStr := make([]string, len(res))
	for i, v := range res {
		resStr[i] = strconv.Itoa(v)
	}
	bw.WriteString(strconv.Itoa(len(resStr)) + "\n" + strings.Join(resStr, " "))
	return nil
}

func getMinPrimes(n int) []int {
	isPrime := newSieve(MaxN)
	if isPrime[n] {
		return []int{n}
	}

	if n%2 == 0 {
		for i := 2; i <= n/2; i++ {
			if isPrime[i] && isPrime[n-i] {
				return []int{i, n - i}
			}
		}
	}

	if isPrime[n-2] {
		return []int{2, n - 2}
	}

	remainder := n - 3
	for i := 2; i <= remainder/2; i++ {
		if isPrime[i] && isPrime[remainder-i] {
			return []int{3, i, remainder - i}
		}
	}

	return nil
}

// зарание вычисляем
func newSieve(max int) []bool {
	isPrime := make([]bool, max+1)
	for i := 2; i <= max; i++ {
		isPrime[i] = true
	}
	for p := 2; p*p <= max; p++ {
		if isPrime[p] {
			for i := p * p; i <= max; i += p {
				isPrime[i] = false
			}
		}
	}
	return isPrime
}

// func isPrime(n int) bool {
// 	if n <= 1 {
// 		return false
// 	}
// 	for i := 2; i*i <= n; i++ {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }
