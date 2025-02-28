package main

import "fmt"

func ConcatAlternate(slice1, slice2 []int) []int {

	if len(slice1) == 0 && len(slice2) > 0 {
		return slice2
	}

	if len(slice2) == 0 && len(slice1) > 0 {
		return slice1
	} 

	newSlice := []int{}
	if len(slice1) == len(slice2) {
		size := len(slice1)
		for i := 0; i < size; i++ {
			newSlice = append(newSlice, slice1[i])
			newSlice = append(newSlice, slice2[i])
		}
	}

	if len(slice1) > len(slice2) {
		size := len(slice2)
		for i := 0; i < size; i++ {
			newSlice = append(newSlice, slice1[i])
			newSlice = append(newSlice, slice2[i])
		}
		newSlice = append(newSlice, slice1[size:]...)
	}

	if len(slice1) < len(slice2) {
		size := len(slice1)
		for i := 0; i < size; i++ {
			newSlice = append(newSlice, slice2[i])
			newSlice = append(newSlice, slice1[i])
		}
		newSlice = append(newSlice, slice2[size:]...)
	}
	return newSlice
}

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}
