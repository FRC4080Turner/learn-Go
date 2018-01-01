package main

import "fmt"

var x int

func increment() int {
	x++
	return x
}

func main() {
	fmt.Println(increment())
	fmt.Println(x)
	fmt.Println(increment())
	fmt.Println(x)
	// boath increment and main have access to x
}
