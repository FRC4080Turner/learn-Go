package main

import "fmt"

func main() {
	fmt.Println(spliter(1))
	fmt.Println(spliter(2))

	half := func(i int) (int, bool) {
		var even bool

		if i%2 == 0 {
			even = true
		} else {
			even = false
		}

		return i / 2, even
	}

	fmt.Println(half(1))
	fmt.Println(half(2))
}

func spliter(i int) (int, bool) {

	var even bool

	if i%2 == 0 {
		even = true
	} else {
		even = false
	}

	return i / 2, even
}
