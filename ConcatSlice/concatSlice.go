package main

import "fmt"

func ConcatSlice(slice1, slice2 []int) []int {
	if len(slice1) == 0 && len(slice2) > 0 {
		return slice2
	}

	if len(slice2) == 0 && len(slice1) > 0 {
		return slice1
	}

	slice1 = append(slice1, slice2...)
	return slice1
}

func main() {
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{}))
}
