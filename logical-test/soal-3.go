package main

import "fmt"

func main() {
	n := 9
	a, b := 0, 1

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a)
		a, b = b, a+b
	}
}
