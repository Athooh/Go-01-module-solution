package main

import "fmt"

func FifthAndSkip(str string) string {
	if str == "" {
		return "\n"
	}

	if len(str) < 5 {
		return "Invalid Input" + "\n"
	}

	fmtStr := ""

	for _, c := range str {
		if c != ' ' {
			fmtStr += string(c)
		}
	}

	nwStr := ""
	count := 0
	for i := 0; i < len(fmtStr); i++ {
		if count == 5 {
			nwStr += string(' ')
			count = 0
			i++
		} 
		nwStr += string(fmtStr[i])
		count++
	}
	return nwStr + "\n"
}

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
}
