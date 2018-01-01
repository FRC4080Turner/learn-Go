package main

import "fmt"

func factorial(x int) int {
	if x < 3 {
		return 2
	}
	return x * factorial(x-1)
}

func main() {
	fmt.Println(factorial(4))
}
