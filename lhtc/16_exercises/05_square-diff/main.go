// Solvs Project Euler question 6 "https://projecteuler.net/problem=6".

package main

import (
	"fmt"
)

func main() {

	var sqOfSum float64
	var sumOfSq float64

	sq := func(x float64) float64 {
		return x * x
	}

	for i := 1.0; i <= 100; i++ {

		sumOfSq += sq(i)
		sqOfSum += i
	}
	sqOfSum = sq(sqOfSum)

	fmt.Println("The difference between the sum of the squares and and the square of the sum:\n", sqOfSum-sumOfSq)
}
