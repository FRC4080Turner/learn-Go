package main

import "fmt"

func main() {

	myFriendsName := "Al"
	greeting := "Wassup"

	switch {
	case len(myFriendsName) == 2:
		fmt.Println(greeting + " my friend with name of length 2")
	case myFriendsName == "Al":
		fmt.Println(greeting + " Al")
	case myFriendsName == "Jenny":
		fmt.Println(greeting + " Jenny")
	}
}
