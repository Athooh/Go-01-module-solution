package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	if len(os.Args) <= 1 {
		return
	}

	args := os.Args[1:]

	for _, str := range args {
		s := ""
		for _, char := range str {
			if char >= 'A' && char <= 'Z' {
				s += string(LowerCase(char))
			} else {
				s += string(char)
			}
		}

		srune := []rune{}

		for _, char := range s {
			srune = append(srune, char)
		}

		for i, c := range srune {
			if i != 0 && c == ' ' && srune[i-1] >= 'a' && srune[i-1] <= 'z' {
				srune[i-1] = UpperCase(srune[i-1])
			} else if i == len(srune)-1 && c >= 'a' && c <= 'z' {
				srune[i] = UpperCase(srune[i])
			}
		}

		st := ""
		for _, char := range srune {
			st += string(char)
		}

		for _, c := range st {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
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

func UpperCase(s rune) rune {
	return rune(s - 32)
}

func LowerCase(s rune) rune {
	return rune(s + 32)
}
