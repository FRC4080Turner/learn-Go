// Print the txt of each line that appers
// more than once in stdin, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)      // make a map with strings as keys that link int values
	input := bufio.NewScanner(os.Stdin) // create a scanner that reads Stdin
	for input.Scan() {                  // for everything the scanner read...
		counts[input.Text()]++ // add the current text scanned and increment the value.
	}
	// ignore potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
