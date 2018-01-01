package main

import "fmt"

func main() {

	age := 21
	fmt.Println(&age) // print the mem address

	changeMe(&age)

	fmt.Println(&age) // address
	fmt.Println(age)  // value
}

func changeMe(z *int) {
	fmt.Println(z)  // printing the mem address
	fmt.Println(*z) // printing the value at an address
	*z = 24         // value at address "z" should be equal to 24
	fmt.Println(z)  // address where "age" is
	fmt.Println(&z) // address were "z" is ?
	fmt.Println(*z) // value at address
}
