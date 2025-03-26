package main 

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
		return
	}

	args := os.Args[1:]
	optionsInt := uint32(0)

	for _, arg := range args {
		// check length of arguements if 1
		if len(arg) == 1 {
			fmt.Println("Invalid Option")
			return
		}
		// check if first character of the arguements is a -
		if len(arg) > 0 && arg[0] == '-' {
			// check is first & second characters of the aguements are -h
			if len(arg) > 1 && arg[:2] == "-h" {
				fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
				return
			}
		} else {
			fmt.Println("Invalid Option")
			return
		}

		for _, ch := range arg[1:] {
			if ch >= 'a' && ch <= 'z' {
				bitPos := ch - 'a'
				optionsInt |= (1 << bitPos)
			} else {
				fmt.Println("Invalid Option")
				return
			}
		}

	}
	bitStr := fmt.Sprintf("%032b", optionsInt)
	
	for i := 0; i < 32; i += 8 {
		if (i + 8) < 32 {
			fmt.Print(bitStr[i:i+8] + " ")
		} else {
			fmt.Print(bitStr[i:i+8])
		}
	}
	fmt.Println()
}