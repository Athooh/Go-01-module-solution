package main

import (
	"fmt"
	"strconv"
)

func NotDecimal(dec string) string {
	// Check if the string is empty
	if dec == "" {
		return "\n"
	}

	// Find the index of the decimal point
	pointIndex := -1
	for i, ch := range dec {
		if ch == '.' {
			pointIndex = i
			break
		}
	}

	// If there is no decimal point, return the original string with a newline
	if pointIndex == -1 {
		return dec + "\n"
	}

	// Split into integer and fractional parts
	intPart := dec[:pointIndex]
	fracPart := dec[pointIndex+1:]

	// If the fractional part is just "0", return the original string
	if fracPart == "0" {
		return dec + "\n"
	}

	// Combine integer and fractional parts to form a whole number
	combined := intPart + fracPart

	// Ensure the combined number is a valid integer
	_, err := strconv.Atoi(combined)
	if err != nil {
		return dec + "\n"
	}

	return combined + "\n"
}

func main() {
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
}
