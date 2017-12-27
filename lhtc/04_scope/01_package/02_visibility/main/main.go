package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	foo()
}

func foo() {
	// won't run because x is declared within main
	fmt.Println(x)
}
