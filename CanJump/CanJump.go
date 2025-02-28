package main

import "fmt"

func CanJump(nums []uint) bool {

	if len(nums) == 0 {
		return false
	}

	if len(nums) == 1 {
		return true
	}

	start := 0

	for start < len(nums)-1 {
		step := int(nums[start])
		if step == 0 || start+step > len(nums) {
			return false
		}
		start += step
	}
	return start == len(nums)-1
}

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2))

	input3 := []uint{0}
	fmt.Println(CanJump(input3))
}
