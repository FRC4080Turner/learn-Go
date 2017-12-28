package main

import "fmt"

func zero(x *int) {
	*x = 0
	fmt.Printf("%p\n", &x)
}

func main() {
	x := 5
	fmt.Println(x)
	zero(&x)
	fmt.Printf("%p\n", &x)
	fmt.Println(x)
}
