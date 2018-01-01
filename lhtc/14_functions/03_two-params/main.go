package main

import "fmt"

func main() {
	greet("Jane", "Doe")
	altGreet("Jane", "Doe")
}

func greet(fname string, lname string) {
	fmt.Println(fname, lname)
}

func altGreet(fname, lname string) { // parameters of same type
	fmt.Println(fname, lname)
}
