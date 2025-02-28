package main

import "fmt"

func WordFlip(str string) string {
	if str == "" {
		return "Invalid Output\n"
	}

	s := SplitwhiteSpace(str)
	revWord := []string{}
	word := ""

	for i := len(s) - 1; i >= 0; i-- {
		revWord = append(revWord, s[i])
	}

	for i, c := range revWord {
		if i < len(revWord)-1 {
			word += c + " "
		} else {
			word += c
		}
	}

	return word + "\n"
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

func main() {
	fmt.Print(WordFlip("First second last"))
	fmt.Print(WordFlip(""))
	fmt.Print(WordFlip("     "))
	fmt.Print(WordFlip(" hello  all  of  you! "))
}
