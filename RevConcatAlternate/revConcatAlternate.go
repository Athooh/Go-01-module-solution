package main

import "fmt"

func RevConcatAlternate(slice1, slice2 []int) []int {

	newSlice := []int{}
	if len(slice1) >= 1 && len(slice2) == 0 {
		size := len(slice1) - 1
		for i := size; i >= 0; i-- {
			newSlice = append(newSlice, slice1[i])
		}
		return newSlice
	}

	if len(slice1) == 0 && len(slice2) >= 1 {
		size := len(slice2) - 1
		for i := size; i >= 0; i-- {
			newSlice = append(newSlice, slice2[i])
		}
		return newSlice
	}

	if len(slice1) == len(slice2) {
		size := len(slice1) - 1
		for i := size; i >= 0; i-- {
			newSlice = append(newSlice, slice1[i])
			newSlice = append(newSlice, slice2[i])
		}
		return newSlice
	}

	if len(slice1) > len(slice2) {
		size1 := (len(slice1) - len(slice2)) - 1
		for i := len(slice1) - 1; i >= size1; i-- {
			newSlice = append(newSlice, slice1[i])
		}

		size := len(slice2) - 1
		for i := size; i >= 0; i-- {
			newSlice = append(newSlice, slice1[i])
			newSlice = append(newSlice, slice2[i])
		}
		return newSlice
	}

	if len(slice1) < len(slice2) {
		size1 := (len(slice2) - len(slice1)) - 1
		for i := len(slice2) - 1; i >= 0; i-- {
			if i == size1 {
				break
			}
			newSlice = append(newSlice, slice2[i])
		}

		size := len(slice1) - 1
		for i := size; i >= 0; i-- {
			newSlice = append(newSlice, slice1[i])
			newSlice = append(newSlice, slice2[i])
		}
		return newSlice
	}
	return newSlice
}

func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
}
