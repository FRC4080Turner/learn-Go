// anonymous function & assigning to variable
package main

import "fmt"

var x = 0

func increment() int {
	x++
	return x
}

func main() {
	fmt.Println(increment())
	fmt.Println(increment())

	y := x
	deincrement := func() int {
		y--
		return y
	}
	fmt.Println(deincrement())
	fmt.Println(deincrement())
}
