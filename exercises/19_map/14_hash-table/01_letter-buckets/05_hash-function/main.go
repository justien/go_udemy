package main

import "fmt"

/*
apply hashBucket to a word, and print the result
*/
func main() {
	n := hashBucket("Go", 12)
	fmt.Println(n)
}

/*
using the word and the number of buckets
take the 1st letter f the word mod number of buckets
and return the specific bucket that word belongs to.
*/
func hashBucket(word string, buckets int) int {
	letter := int(word[0])
	myBucket := letter % buckets
	return myBucket
}
