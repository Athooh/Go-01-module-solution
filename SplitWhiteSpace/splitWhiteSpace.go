package main

import "fmt"

func SplitwhiteSpace(s string) []string {
	var words []string
	var word string
	for _, char := range s {
		if char == ' ' {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(char)
		}
	}
	if word != "" {
		words = append(words, word)
	}
	return words
}

func main() {
	s := " Hello World how are you "
	words := SplitwhiteSpace(s)
	fmt.Println(words)
}
