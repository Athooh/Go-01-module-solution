package main

import "fmt"

func CanJump(num []uint) bool {
	if len(num) == 0 {
		return false
	}

	if len(num) == 1 {
		return true
	}

	start := 0

	for start < len(num)-1 {
		step := int(num[start])

		if step == 0 || start+step > len(num) {
			return false
		}
		start += step
	}
	return start == len(num)-1
}

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2))

	input3 := []uint{0}
	fmt.Println(CanJump(input3))
}
