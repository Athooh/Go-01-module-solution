package main

import "fmt"

func ThirdTimeIsACharm(str string) string {
	count := 1
	thirdChar := ""

	if len(str) < 3 || str == "" {
		return "\n"
	}

	for _, char := range str {
		if count == 3 {
			thirdChar += string(char)
			count = 1
		} else {
			count++
		}
	}
	return thirdChar + "\n"
}

func main() {
	fmt.Print(ThirdTimeIsACharm("123456789"))
	fmt.Print(ThirdTimeIsACharm(""))
	fmt.Print(ThirdTimeIsACharm("a b c d e f"))
	fmt.Print(ThirdTimeIsACharm("12"))
}
