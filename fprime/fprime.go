package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	number, err := Atoi(os.Args[1])
	if !err || number <= 1{
		return
	}

	factors := Primefactor(number)

	fprimeStr := ""

	for i, n := range factors {
		if i < len(factors) - 1 {
			fprimeStr += Itoa(n) + "*"
		} else {
			fprimeStr += Itoa(n)
		}
	}

	for _, c := range fprimeStr {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func Primefactor(n int) []int {
	factors := []int{}
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}
	return factors
}

func Atoi(s string) (int, bool) {
	sign := 1
	number := 0

	for i, c := range s {
		if i == 0 && c == '-' {
			sign = -1
		} else if i == 0 && c == '+' {
			sign = 1
		} else if c >= '0' && c <= '9' {
			number = number * 10 + int(c - '0')
		} else {
			return 0, false
		}
	}
	return (sign * number), true
}

func Itoa(n int) string {
	sign := ""

	if n < 0 {
		n = -n
		sign = "-"
	}

	if n == 0 {
		return "0"
	}

	digits := []rune{}

	for n > 0 {
		digit := n % 10
		digits = append([]rune{rune('0' + digit)}, digits...)
		n /= 10
	}
	return sign + string(digits)
}
