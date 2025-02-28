package main 

import "fmt"

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

func main() {
    fmt.Println(Itoa(12345))
    fmt.Println(Itoa(0))
    fmt.Println(Itoa(-1234))
    fmt.Println(Itoa(987654321))
}