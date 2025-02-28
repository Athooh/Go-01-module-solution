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
	piglatin := ""

	for i, c := range str {
		if i == 0 && IsVowels(c) {
			piglatin = str + "ay"
		} else if i != 0 && IsVowels(c) {
			piglatin = str[i:] + str[:i] + "ay"
		}
	}

	vowel := false
	for _, c := range str {
		if IsVowels(c) {
			vowel = true
			break
		}
	}

	if !vowel {
		piglatin = "No vowels"
	}

	for _, c := range piglatin {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func IsVowels(s rune) bool {
	return s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' || s == 'A' || s == 'E' || s == 'I' || s == 'O' || s == 'U'
}
