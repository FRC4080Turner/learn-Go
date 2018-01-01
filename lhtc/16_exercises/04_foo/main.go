package main

import "fmt"

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4, 5}
	foo(aSlice...)
	foo()
	fmt.Println("done")
}

func foo(list ...int) {
	fmt.Println(len(list))
}
