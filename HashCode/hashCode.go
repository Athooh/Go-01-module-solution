package main

import "fmt"

func HashCode(dec string) string {
	var hash string
	for _, char := range dec {
		resultHash := (int(char) + len(dec)) % 127
		if resultHash < 32 || resultHash > 126 {
			resultHash += 33
		}
		hash += string(rune(resultHash))
	}
	return hash
}

func main() {
	fmt.Println(HashCode("A"))           // B
	fmt.Println(HashCode("AB"))          // CD
	fmt.Println(HashCode("BAC"))         // EDF
	fmt.Println(HashCode("Hello World")) // Spwwz+bz}wo
}
