package main

import "fmt"

func main() {
	greeting := "Wassup"
	switch "Marcus" {
	case "Tim":
		fmt.Println(greeting + " Tim")
	case "Jenny":
		fmt.Println(greeting + " Jenny")
	case "Marcus":
		fmt.Println(greeting + " Marcus")
		fallthrough // run the next automaticaly if you get here
	case "Medhi":
		fmt.Println(greeting + " Medhi")
		fallthrough
	case "Julian":
		fmt.Println(greeting + " Julian")
	case "Sushant":
		fmt.Println(greeting + " Sushant")
	}
}
