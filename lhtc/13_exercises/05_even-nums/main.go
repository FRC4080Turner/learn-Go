package main

import "fmt"

func main() {

	for i := 0; 100 >= i; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
