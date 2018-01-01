package main

import "fmt"

func main() {

	greatest := greatest(22.5, 22.4, 10, 5.5, 0, 2)

	fmt.Println(greatest)

}

func greatest(list ...float64) float64 {

	var greatest float64

	for _, num := range list {

		if num > greatest {
			greatest = num
		}

	}

	return greatest
}
