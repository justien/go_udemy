package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	response, error_message := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")

	if error_message != nil {
		log.Fatal(error_message)
	}

	/*
	I don't get why error_message appears again
	*/

	slice_O_bytes, error_message := ioutil.ReadAll(response.Body)
	response.Body.Close()

	if error_message != nil {
		log.Fatal(error_message)
	}

	fmt.Printf("%s", slice_O_bytes)

}
