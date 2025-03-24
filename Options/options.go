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
	// fmt.Println(args)

	var optionsInt uint32 = 0

	for _, arg := range args {
		if len(arg) > 0 && arg[0] == '-' {
			if len(arg) == 1 {
				fmt.Println("Invalid Option")
				return
			}
			
			if len(arg) > 1 && arg[:2] == "-h" {
				fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
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
		} else {
			fmt.Println("Invalid Option")
			return
		}
	}

	bitStr := fmt.Sprintf("%032b", optionsInt)


	for  i := 0; i < 32; i+= 8 {
		if (i + 8) < 32 {
			fmt.Print(bitStr[i:i+8] + " ")
		} else {
			fmt.Print(bitStr[i:i+8])
		}
	}
	fmt.Println()
}