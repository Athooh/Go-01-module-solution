package main 

import "fmt"

func Atoi( s string) int {
	var n int
	sign := 1

	for i, char := range s {
		if char == '-' && i == 0 {
			sign = -1
		} else if char == '+' && i == 0 {
			sign = 1
		} else if char >= '0' && char <= '9' {
			n = n * 10 + int(char - '0')
		} else {
			return 0
		}
	}
	return sign * n
}

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}