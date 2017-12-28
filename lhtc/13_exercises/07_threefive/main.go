package main

import (
	"fmt"
)

func main() {

	output := 0

	for i := 0; i <= 1000; i++ {
		if i%3 == 0 {
			output += i
		} else if i%5 == 0 {
			output += i
		}
	}
	fmt.Println("Answer is:", output)
}
