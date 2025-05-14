package main

import (
	"fmt"
	"unicode"
)

func hitungAngka(input []string) int {
	hitung := 0
	for _, char := range input {
		if unicode.IsDigit(rune(char[0])) {
			hitung++
		}
	}
	return hitung
}
func main() {
	testCases := [][]string{
		{"2", "h", "6", "u", "y", "t", "7", "j", "y", "h", "8"},
		{"b", "7", "h", "6", "h", "k", "i", "5", "g", "7", "8"},
		{"7", "b", "8", "5", "6", "9", "n", "f", "y", "6", "9"},
		{"u", "h", "b", "n", "7", "6", "5", "1", "g", "7", "9"},
	}

	for i, test := range testCases {
		fmt.Printf("Soal %d: %v\n", i+1, test)
		fmt.Printf("Output: %d\n\n", hitungAngka(test))
	}
}
