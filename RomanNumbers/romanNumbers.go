// package main

// import (
// 	"os"

// 	"github.com/01-edu/z01"
// )

// func rn(n int) string {

// 	var romanNumbers []string
// 	var romanMethod []string

// 	for n > 0 {
// 		if n >= 1000 {
// 			romanNumbers, romanMethod = append(romanNumbers, "M"), append(romanMethod, "M")
// 			n -= 1000
// 		} else if n >= 900 {
// 			romanNumbers, romanMethod = append(romanNumbers, "CM"), append(romanMethod, "(M-C)")
// 			n -= 900
// 		} else if n >= 500 {
// 			romanNumbers, romanMethod = append(romanNumbers, "D"), append(romanMethod, "D")
// 			n -= 500
// 		} else if n >= 400 {
// 			romanNumbers, romanMethod = append(romanNumbers, "CD"), append(romanMethod, "(D-C)")
// 			n -= 400
// 		} else if n >= 100 {
// 			romanNumbers, romanMethod = append(romanNumbers, "C"), append(romanMethod, "C")
// 			n -= 100
// 		} else if n >= 90 {
// 			romanNumbers, romanMethod = append(romanNumbers, "XC"), append(romanMethod, "(C-X)")
// 			n -= 90
// 		} else if n >= 50 {
// 			romanNumbers, romanMethod = append(romanNumbers, "L"), append(romanMethod, "L")
// 			n -= 50
// 		} else if n >= 40 {
// 			romanNumbers, romanMethod = append(romanNumbers, "XL"), append(romanMethod, "(L-X)")
// 			n -= 40
// 		} else if n >= 10 {
// 			romanNumbers, romanMethod = append(romanNumbers, "X"), append(romanMethod, "X")
// 			n -= 10
// 		} else if n >= 9 {
// 			romanNumbers, romanMethod = append(romanNumbers, "IX"), append(romanMethod, "(X-I)")
// 			n -= 9
// 		} else if n >= 5 {
// 			romanNumbers, romanMethod = append(romanNumbers, "V"), append(romanMethod, "V")
// 			n -= 5
// 		} else if n >= 4 {
// 			romanNumbers, romanMethod = append(romanNumbers, "IV"), append(romanMethod, "(V-I)")
// 			n -= 4
// 		} else if n >= 1 {
// 			romanNumbers, romanMethod = append(romanNumbers, "I"), append(romanMethod, "I")
// 			n -= 1
// 		}
// 	}

// 	rM := ""
// 	rN := ""
	
// 	for i, c := range romanMethod {
// 		if i < len(romanMethod) - 1 {
// 			rM += c + "+"
// 		} else {
// 			rM += c
// 		}
// 	}

// 	for _, c := range romanNumbers {
// 		rN += c
// 	}

// 	return rM + "\n" + rN
// }

// func main() {
// 	if len(os.Args) != 2 {
// 		return
// 	}

// 	er := "ERROR: cannot convert to roman digit"

// 	n := os.Args[1]

// 	numb, err := Atoi(n) 
// 	if !err {
// 		for _, c := range er {
// 			z01.PrintRune(c)
// 		}
// 		z01.PrintRune('\n')
// 		return
// 	}

// 	if numb >= 4000 || numb <= 0 {
// 		for _, c := range er {
// 			z01.PrintRune(c)
// 		}
// 		z01.PrintRune('\n')
// 		return
// 	}

// 	result := rn(numb)
// 	for _, c := range result {
// 		z01.PrintRune(c)
// 	}
// 	z01.PrintRune('\n')
// }

// func Atoi(s string) (int, bool) {
// 	sign := 1
// 	num := 0
// 	for i, ch := range s {
// 		if i == 0 && ch == '-' {
// 			sign = -1
// 		} else if i == 0 && ch == '+' {
// 			sign = 1
// 		} else if ch >= '0' && ch <= '9' {
// 			num = num * 10 + int(ch - '0')
// 		} else {
// 			return 0, false
// 		}
// 	}
// 	return sign * num, true
// }

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func rn(n int) string {
	var romanNumbers []string
	var romanMethod []string

	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	methods := []string{"M", "(M-C)", "D", "(D-C)", "C", "(C-X)", "L", "(L-X)", "X", "(X-I)", "V", "(V-I)", "I"}

	for i, val := range values {
		for n >= val {
			romanNumbers = append(romanNumbers, symbols[i])
			romanMethod = append(romanMethod, methods[i])
			n -= val
		}
	}

	// Join the calculation method string
	rM := ""
	for i, c := range romanMethod {
		if i != 0 {
			rM += "+"
		}
		rM += c
	}

	// Join the final Roman numeral string
	rN := ""
	for _, c := range romanNumbers {
		rN += c
	}

	return rM + "\n" + rN
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	errorMsg := "ERROR: cannot convert to roman digit"

	n := os.Args[1]

	numb, valid := Atoi(n)
	if !valid || numb <= 0 || numb >= 4000 {
		for _, c := range errorMsg {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
		return
	}

	result := rn(numb)
	for _, c := range result {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func Atoi(s string) (int, bool) {
	sign := 1
	num := 0
	for i, ch := range s {
		if i == 0 && ch == '-' {
			sign = -1
		} else if i == 0 && ch == '+' {
			sign = 1
		} else if ch >= '0' && ch <= '9' {
			num = num*10 + int(ch-'0')
		} else {
			return 0, false
		}
	}
	return sign * num, true
}
