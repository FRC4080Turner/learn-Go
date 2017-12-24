// Echo prints its command-line args.

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Assign multiple variables as strings
	var arg string
	separator := "\n"

	/*
		Set i to one, and as long as it is
		less than the length of arguments
		keep running.
	*/
	for i := 1; i < len(os.Args); i++ {
		// Arguments on separate line w/indicators
		arg += "Argument #" + strconv.Itoa(i) + ": " + os.Args[i] + separator

	}
	fmt.Println(arg)
}
