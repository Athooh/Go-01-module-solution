package main 

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {

	str1 := os.Args[1]
	str2 := os.Args[2]

	var idx1 int
	var idx2 int

	for idx1 < len(str1) && idx2 < len(str2) {
		if str1[idx1] == str2[idx2] {
			idx1++
		}
		idx2++
	}

	if idx1 == len(str1) {
		for _, c := range str1 {
			z01.PrintRune(c)
		}
	}
	z01.PrintRune('\n')
}