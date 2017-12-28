package main

import "fmt"

func main() {
	for i := 1000; i <= 1100; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
	foo := 'a'
	var x uint32 = 32
	fmt.Println(foo)
	fmt.Printf("%T\n", foo)
	fmt.Printf("%T\n", rune(foo))
	fmt.Printf("%#x\n", foo)
	fmt.Printf("%#x\n", x)
}
