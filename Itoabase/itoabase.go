package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	for i := 0; i < 30; i++ {
		value := random.IntBetween(-1000000, 1000000)
		base := random.IntBetween(2, 16)
		challenge.Function("ItoaBase", ItoaBase, solutions.ItoaBase, value, base)
	}
	for i := 0; i < 5; i++ {
		base := random.IntBetween(2, 16)
		challenge.Function("ItoaBase", ItoaBase, solutions.ItoaBase, random.MaxInt, base)
		challenge.Function("ItoaBase", ItoaBase, solutions.ItoaBase, random.MinInt, base)
	}
}

func ItoaBase(value, base int) string {
	strChar := "0123456789ABCDEF"

	result := ""
	sign := ""

	if value < 0 {
		value = -value
		sign = "-"
	}

	if value == 0 {
		return "0"
	}

	n := uint(value)
	b := uint(base)

	for n > 0 {
		remainder := n % b
		result = string(strChar[remainder]) + result
		n /= b
	}
	return sign + result
}
