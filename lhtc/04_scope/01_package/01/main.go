package main

import "fmt"

var x int32 = 42

func main() {
	fmt.Println(x)
	y := 17
	foo()
	fmt.Println(y)
}

func foo() {
	fmt.Println(x)
}
