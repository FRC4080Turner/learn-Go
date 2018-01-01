package main

import "fmt"

func main() {

	name := "Todd"
	fmt.Println(&name) // print the mem address

	changeMe(&name)

	fmt.Println(&name) // address
	fmt.Println(name)  // value
}

func changeMe(z *string) {
	fmt.Println(z)  // printing the mem address
	fmt.Println(*z) // printing the value at an address
	*z = "Rocky"    // value at address "z" should be equal to 24
	fmt.Println(z)  // address where "name" is
	fmt.Println(&z) // address were "z" is ?
	fmt.Println(*z) // value at address
}
