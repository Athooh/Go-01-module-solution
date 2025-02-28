package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] == "" {
		z01.PrintRune('\n')
		return
	}

	args := SplitWhiteSpaces(os.Args[1])
	str := ""

	for i, char := range args {
		if i < len(args)-1 {
			str += char + " "
		} else {
			str += char
		}
	}

	for _, c := range str {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func SplitWhiteSpaces(s string) []string {
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
