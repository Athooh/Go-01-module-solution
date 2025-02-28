package main 

import "fmt"

func ZipString(s string) string {

	count := 1
	zip := ""

	for i := 1; i <= len(s); i++ {
		if i < len(s) && s[i] == s[i-1] {
			count++
		} else {
			zip += Itoa(count) + string(s[i-1])
			count = 1
		}
	}
	return zip
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

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}