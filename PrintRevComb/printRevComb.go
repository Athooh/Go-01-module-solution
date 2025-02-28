package main

import "github.com/01-edu/z01"

func main() {
	first := true // Track the first combination to avoid a leading comma

	for i := '9'; i >= '2'; i-- {
		for j := i - 1; j >= '1'; j-- {
			for k := j - 1; k >= '0'; k-- {
				// Ensure first digit > second digit > third digit
				if i > j && j > k {
					// Print comma and space except for the first output
					if !first {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
					// Print the combination
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
					first = false
				}
			}
		}
	}
	z01.PrintRune('\n') // End with a newline
}
