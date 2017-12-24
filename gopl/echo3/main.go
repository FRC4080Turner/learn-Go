// Improved Echo
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[:], ", "))
	fmt.Println("\n" + strings.Join(os.Args[1:], ", "))
	// strings.Join saparates string array elements with the specified separator
}
