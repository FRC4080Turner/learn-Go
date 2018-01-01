package main

import "fmt"

func main() {
	greet("Jane") // Pass in an argument...
	greet("John")
}

func greet(name string) { // to set a parameter.
	fmt.Println(name)

}
