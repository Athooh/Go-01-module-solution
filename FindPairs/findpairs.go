package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findPairs(arr []int, target int) [][]int{
	result := [][]int{}
	
	for i := 0; i < len(arr); i++ {
		for j := i+1; j < len(arr); j++ {
			if arr[i] + arr[j] == target {
				pair := []int{i, j}
				result = append(result, pair)
			}
		}
	}
	return result
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid input.")
		return
	}

	args1 := os.Args[1]
	args2 := os.Args[2]

	arr := []int{}

	if !strings.HasPrefix(args1, "[") || !strings.HasSuffix(args1, "]") {
		fmt.Println("Invalid input.")
		return
	}

	target, err := strconv.Atoi(args2)
	if err != nil {
		fmt.Println("Invalid target sum.")
		return
	}

	// fmt.Println(target)

	args1 = args1[1:len(args1)-1]
	// fmt.Println(args1)

	arrStr := strings.Split(args1, ", ")

	// fmt.Println(arrStr)

	for _, v := range arrStr {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("Invalid number: %s\n", v)
			return
		}
		arr = append(arr, num)
	}
	// fmt.Println(arr)

	result := findPairs(arr, target)
	if len(result) == 0 {
		fmt.Println("No pairs found.")
		return
	} else {
	fmt.Printf("Pairs with sum %d: %v\n", target, result)
	}
}

