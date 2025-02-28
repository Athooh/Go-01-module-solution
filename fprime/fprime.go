package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n == 0 {
		return
	}

	factors := PrimeFactor(n)
	if len(factors) == 0 {
		return
	}

	for i, factor := range factors {
		for _, num := range Itoa(factor) {
			z01.PrintRune(num)
		}
		if i < len(factors) -1 {
			z01.PrintRune('*')
		}
	}
	z01.PrintRune('\n')
}


func PrimeFactor(n int) []int {
	factors := []int{}
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}
	return factors
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

	var digits []rune

	for n > 0 {
		digit := n%10
		digits = append([]rune{rune('0' + digit)}, digits...)
		n /= 10
	}
	return sign + string(digits)
}