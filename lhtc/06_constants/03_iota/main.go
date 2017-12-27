package main

import "fmt"

const (
	// A multi init
	A = iota // 0
	B        // B = 1
	C        // C = 2
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
