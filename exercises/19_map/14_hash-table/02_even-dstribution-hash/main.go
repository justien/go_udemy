package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {

	// get the book adventures of sherlock holmes
	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
	if err != nil {
		log.Fatal(err)
	}

	// scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)

	// Create slice to hold counts
	buckets := make([]int, 12)

	/*
	  For each word, use the hashBucket function to decide
	  which bucket it belongs to.
	  then increment the count associated with that bucket
	*/

	for scanner.Scan() {
		n := hashBucket(scanner.Text(), 12)
		buckets[n]++
	}
	fmt.Println(buckets)
}

/*
add together the byte value of every letter in a word,
divide this by 12, and return the remainder.
(this will skew away from an even distribution depending on
the letter frequency of the input languages)
*/

func hashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}
