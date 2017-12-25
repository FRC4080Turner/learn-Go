// Print the txt of each line that appers
// more than once in stdin, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) // make a map with strings as keys that link int values
	files := os.Args[1:]           // get all args, assume they are file paths
	if len(files) == 0 {           // use stdin if no files were given as args
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			file, error := os.Open(arg) // get the error and file content from the list of args (files)
			if error != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", error)
				continue
			}
			countLines(file, counts)
			file.Close()
		}
	}
	// ignore potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) { // the same map made in main is passed in and edited
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}

}
