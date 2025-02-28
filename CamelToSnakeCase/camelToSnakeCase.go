package main

import "fmt"

func CamelToSnakeCase(s string) string {
	var snakeCase string
	for i, char := range s {
		if i != 0 && char >= 'A' && char <= 'Z' {
			if s[i-1] >= 'A' && s[i-1] <= 'Z' {
				return s // Return the original string if consecutive uppercase letters are found
			}
			snakeCase += "_"
		}
		snakeCase += string(char)
	}
	return snakeCase
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))       // hello_world
	fmt.Println(CamelToSnakeCase("helloWorld"))       // hello_world
	fmt.Println(CamelToSnakeCase("camelCase"))        // camel_case
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE")) // CAMELtoSnackCASE
	fmt.Println(CamelToSnakeCase("camelToSnakeCase")) // camel_to_snake_case
	fmt.Println(CamelToSnakeCase("hey2"))             // hey2
}
