package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {

	// get the book moby dick

	response, error := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if error != nil {
		log.Fatal(error)
	}

	/*
			scan the page
			NewScanner takes a reader and
			response.Body implements the reader interface (so it is a reader)

		  .. oh and btw just before you go, pls close response.Body
	*/
	scanner := bufio.NewScanner(response.Body)
	defer response.Body.Close()

	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)

	// Loop over the words
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
