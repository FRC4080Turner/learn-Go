package main

import "fmt"

func hello() {
	fmt.Print("hello ")
}

func world() {
	fmt.Println("world")
}

func main() {
	fmt.Println("one")
	defer world()
	fmt.Println("two")
	hello()
	fmt.Println("three")
}
