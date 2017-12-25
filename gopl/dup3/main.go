// Print the txt of each line that appers
// more than once in a file.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int) // make a map with strings as keys that link int values
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename) // read the file into the variable "data". ReadFile outputs file as a byte slice
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") { // split the file into its induvidulal lines
			counts[line]++ // add text as a key and increment every time the same text shows up
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
