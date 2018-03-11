package main

import (
	"bufio"
	"log"
	"net/http"
)

func main() {

	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	scanner.Split(bufio.ScanWords)

	buckets := make(map[int]map[string]int)

	for i := 0; i < 12; i++ {
		buckets[i] = make(map[string]int)
	}
}
