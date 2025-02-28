package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}

	str := SplitwhiteSpace(os.Args[1])
	nwStr := []string{}

	nwStr = append(nwStr, str[1:]...)
	nwStr = append(nwStr, str[0])
	rostring := ""

	for i, c := range nwStr {
		if i < len(nwStr)-1 {
			rostring += c + " "
		} else {
			rostring += c
		}
	}
	
	for _, c := range rostring {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func SplitwhiteSpace(s string) []string {
	words := []string{}
	word := ""

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
