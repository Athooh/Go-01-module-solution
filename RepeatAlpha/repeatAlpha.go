package main 

import "fmt"

func RepeatAlpha(s string) string {
	var repeatWord string

	for _, char := range s {
		var repCount int
		if char >= 'a' && char <= 'z' {
			repCount = int(char - 'a' + 1)
			for i:= 0; i < repCount; i++ {
				repeatWord += string(char)
			}
		} else if char >= 'A' && char <= 'Z' {
			repCount = int(char - 'A' + 1)
			for i:= 0; i < repCount; i++ {
				repeatWord += string(char)
			}
		} else {
			repeatWord += string(char)
		}
	}
	return repeatWord
}

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}