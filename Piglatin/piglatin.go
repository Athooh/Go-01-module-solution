package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	str := os.Args[1]

	pigLatin := ""

	for i, c := range str {
		if i == 0 && IsVowel(c) {
			pigLatin = str + "ay"
		} else if IsVowel(c) {
			pigLatin = str[i:] + str[:i] + "ay"
		}
	}

	count := 0 

	for _, c := range str {
		if IsVowel(c) {
			break
		} else {
			count++
		}
	}

	if len(str) == count {
		pigLatin = "No Vowels"
	}

	for _, c := range pigLatin {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func IsVowel(c rune) bool {
	return (c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U')
}