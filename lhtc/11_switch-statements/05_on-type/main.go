// Shows switching by variable type.
package main

import "fmt"

// Contact is a test coustom type.
type Contact struct {
	greeting string
	name     string
}

func switchOnType(x interface{}) {
	switch x.(type) { // assert x is of type in the case condition
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("Contact")
	default:
		fmt.Println("unknown")
	}

}

func main() {
	switchOnType(7)
	switchOnType("McLeod")
	var t = Contact{"Good to see you, ", "Tim"}
	switchOnType(t)
}
