package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {

	if len(os.Args) != 2 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	if n <= 0 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	primeSum := 0

	for i:= 2; i <= n; i++ {
		if IsPrime(i) {
			primeSum += i
		}
	}

	p := Itoa(primeSum)
	for _, c := range p {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i:= 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}


func Itoa(n int) string {
	sign := ""
	if n == 0 {
		return "0"
	}

	if n < 0 {
		sign = "-"
		n = -n
	}

	var digits []rune

	for n > 0 {
		digit := n%10
		digits = append([]rune{rune('0' + digit)}, digits...)
		n /= 10
	}

	return sign + string(digits)
}