package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	response, error := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if error != nil {
		log.Fatalln(error)
	}

	mybytes, _ := ioutil.ReadAll(response.Body)
	str := string(mybytes)
	fmt.Println(str)

}
