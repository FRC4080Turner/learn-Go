package main

import "fmt"

func main() {
	fmt.Println(spliter(1))
	fmt.Println(spliter(2))

	half := func(i int) (int, bool) {
		return i / 2, i%2 == 0
	}

	fmt.Println(half(1))
	fmt.Println(half(2))
}

func spliter(i int) (int, bool) {
	return i / 2, i%2 == 0
}
