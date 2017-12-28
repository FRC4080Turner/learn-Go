package main

import "fmt"

func main() {
	for i := 2000; i <= 3140; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
}
