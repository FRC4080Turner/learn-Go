package main

import (
	"fmt"
	"strconv"
)

func main() {

	var output string

	for i := 0; i <= 100; i++ {

		// convert "i" to a string and add to output
		output += strconv.Itoa(i)

		if i == 0 {
			// don't add anything to output
			// zero not a multiple of anything
		} else {
			// check for multiple of three
			if i%3 == 0 {
				// add "Fizz to output"
				output += " Fizz"
			}
			// check for multiple of five
			if i%5 == 0 {
				// add "Buzz to the output"
				output += " Buzz"
			}
		}

		fmt.Println(output)

		// reset output
		output = ""
	}
}
