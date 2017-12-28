package main

import "fmt"

func main() {

	// setup variables
	var largeNum int64
	var smallNum int64
	var firstNum int64
	var secondNum int64

	// get user input
	fmt.Print("Enter first number: ")
	fmt.Scan(&firstNum)
	fmt.Print("Enter second number: ")
	fmt.Scan(&secondNum)

	// find the larger of the numbers
	if secondNum >= firstNum {
		largeNum = secondNum
		smallNum = firstNum
	} else {
		largeNum = firstNum
		smallNum = secondNum
	}

	// calculate the remainder
	remainder := largeNum % smallNum

	// show the user the remainder
	fmt.Println("Remainder =", remainder)
}
