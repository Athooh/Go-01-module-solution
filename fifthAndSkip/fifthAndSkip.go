package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"

	// "fmt"
)

func FifthAndSkip(str string) string {
	if str == "" {
		return "\n"
	}

	if len(str) < 5 {
		return "Invalid Input\n"
	}

	frmtStr := ""
	skip := ""

	for _, c := range str {
		if c == ' ' {
			continue
		} else {
			frmtStr += string(byte(c))
		}
	}

	count := 0

	for _, c := range frmtStr {
		count++
		if count == 6{
			skip += " "
			count = 0
		} else {
			skip += string(c)
		}
	}
	return skip + "\n"
}

// func main() {
// 	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
// 	fmt.Print(FifthAndSkip("This is a short sentence"))
// 	fmt.Print(FifthAndSkip("1234"))
// 	fmt.Print(FifthAndSkip("e 5£ @ 8* 7 =56 ;"))
// }


func main() {
	table := []string{"1234556789", "e 5£ @ 8* 7 =56 ;", "QKplq%QSw", "", "hello \\! n4ght cr3a8ure7 ", "Kimetsu no Yaiba", "8595485-52", "-552", "w58tw7474abc", "Po65 4o"}
	for _, s := range table {
		challenge.Function("FifthAndSkip", FifthAndSkip, solutions.FifthAndSkip, s)
	}
}
