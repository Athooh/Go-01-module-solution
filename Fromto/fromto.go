package main

import (
	"fmt"
)

func FromTo(from int, to int) string {
	result := ""

	if from == to {
		return Itoa(from) + "\n"
	}

	if from > 99 || from <= 0 || to > 99 || to <= 0 {
		return "Invalid\n"
	}

	if from < to {
		for i := from; i <= to; i++ {
			if i < 10 {
				result += "0" + Itoa(i) + ", "
			} else if i >= 10 && i < to {
				result += Itoa(i) + ", "
			} else {
				result += Itoa(i)
			}
		}
	} else {
		for i := from; i >= to; i-- {
			if i < 10 {
				if i != 1 {
					result += "0" + Itoa(i) + ", "
				} else {
					result += "0" + Itoa(i)
				}
			} else if i >= from && i >= 10 {
				result += Itoa(i) + ", "
			} else {
				result += Itoa(i)
			}
		}
	}
	return result + "\n"
}

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
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
		digit := n % 10
		digits = append([]rune{rune('0' + digit)}, digits...)
		n /= 10
	}
	return sign + string(digits)
}
