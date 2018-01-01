package main

import "fmt"

func main() {
	fmt.Println(true, true && false || false && true || !(false && false))
}
