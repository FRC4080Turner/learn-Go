package main

import "fmt"

const (
	_ = iota      // 0
	B = iota * 10 // B = 1 * 10
	C = iota * 10 // C = 2 * 10
)

func main() {
	fmt.Println(B)
	fmt.Println(C)
}
