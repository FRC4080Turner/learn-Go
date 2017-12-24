// Echo prints its command-line args.

package main

import (
	"fmt"
	"os"
)

func main() {
	// Assign multiple variables as strings
	var arg, separator string

	/*
		Set i to one, and as long as it is
		less than the length of arguments
		keep running.
	*/
	for i := 1; i < len(os.Args); i++ {
		arg += separator + os.Args[i]
		separator = " "
	}
	fmt.Println(arg)
}
