package main

import "fmt"

func WeAreUnique(str1, str2 string) int {

	if str1 == "" && str2 == "" {
		return -1
	}

	count :=  0

	unique := make(map[rune]bool)
	unseen := make(map[rune]bool)

	for _, char := range str1 {
		unique[char] = true
	}

	for _, char := range str2 {
		unseen[char] = true
	}

	for _, char := range str1 {
		if !unseen[char] {
			count++
		}
	}

	for _, char := range str2 {
		if !unique[char] {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}
