package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]

	unique := make(map[rune]bool)
	unseen := make(map[rune]bool)

	for _, c := range str2 {
		unique[c] = true
	}

	for _, c := range str1 {
		if unique[c] && !unseen[c]{
			unseen[c] = true
			z01.PrintRune(c)
		}

	}
	z01.PrintRune('\n')
}
