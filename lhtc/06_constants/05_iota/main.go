package main

import "fmt"

const (
	A = iota // A = 0
	B        // B = 1
	C        // C = 1
)

// iota count restarts
const (
	D = iota // D = 0
	E        // E = 1
	F        // F = 1
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
}
