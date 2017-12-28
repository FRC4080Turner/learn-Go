package main

import "fmt"

func main() {
	greeting := "Wassup"
	switch "Jenny" {
	case "Tim", "Jenny":
		fmt.Println(greeting + " Tim, or, err, Jenny")
	case "Marcus", "Medhi":
		fmt.Println(greeting + " Both of your names start with 'M'.")
	case "Julian", "Sushant":
		fmt.Println(greeting + " Julian / Sushant")
	}
}
