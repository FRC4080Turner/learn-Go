package main

import "fmt"

func main() {
	
	var x int
	
	imcrement := func() int {
		x++
		return x
	}

	fmt.Println(imcrement())
	fmt.Println(imcrement())
}
