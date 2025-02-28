package main

import "fmt"

func DigitLen(n, base int) int {

	if n < 0 {
		n = -n
	}

	if n == 0 {
		return 1
	}

	if base < 2 || base > 36 {
		return -1
	}

	len := 0

	for n > 0 {
		n /= base
		len++
	}

	return len
}

func main() {
	fmt.Println(DigitLen(100, 10))
	fmt.Println(DigitLen(100, 2))
	fmt.Println(DigitLen(-100, 16))
	fmt.Println(DigitLen(100, -1))
}
