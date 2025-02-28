package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}
	str := SplitWhiteSpaces(s)

	for _, word := range str {
		c := word[0]
		if c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' || c == '!' {
			continue
		} else {
			return false
		}
	}

	return true
}

func SplitWhiteSpaces(s string) []string {
	var words []string
	var word string
	for _, c := range s {
		if c == ' ' {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(c)
		}
	}
	if word != "" {
		words = append(words, word)
	}
	return words
}

func main() {
	table := []string{
		"Hello! €How are you?",
		"a",
		"z",
		"!",
		"Hello How Are 4you",
		"What's this 4?",
		"Whatsthis4",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890",
		"131!",
		"H3110 W0r1d!",
		"",
		" ",
	}

	for _, arg := range table {
		challenge.Function("IsCapitalized", IsCapitalized, solutions.IsCapitalized, arg)
	}
}
