package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		result := fmt.Sprintf("%d", i)

		if i%3 == 0 && i%5 == 0 {
			result += " FizzBuzz"
		} else if i%3 == 0 {
			result += " Fizz"
		} else if i%5 == 0 {
			result += " Buzz"
		}

		fmt.Println(result)
	}
}
