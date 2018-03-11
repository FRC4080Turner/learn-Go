package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

const (
	numberOfBuckets = 12
)

func main() {
	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	scanner.Split(bufio.ScanWords)

	buckets := make([][]string, numberOfBuckets)

	// don't need because the above will already set them to string slices
	/*
		for i := 0; i < numberOfBuckets; i++ {
			buckets[i] = []string{}
		}
	*/

	for scanner.Scan() {
		word := scanner.Text()
		n := hashBucket(word, numberOfBuckets)
		buckets[n] = append(buckets[n], word)
	}

	for i := 0; i < 12; i++ {
		fmt.Println(i, " - ", len(buckets[i]))
	}

	fmt.Println("-----------")
	fmt.Println(len(buckets))
	fmt.Println(cap(buckets))

}

func hashBucket(word string, buckets int) int {
	var sum int
	for _, letter := range word {
		sum += int(letter)
	}
	return sum % buckets
}
