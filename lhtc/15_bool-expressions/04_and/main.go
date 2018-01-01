package main

import "fmt"

func main() {
	if true && false {
		fmt.Println("This didn't run.")
	}

	if true && true {
		fmt.Println("This ran.")
	}
}
