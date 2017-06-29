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

/*

func Get(url string) (resp *Response, err error)
v
... Response is a struct with a bunch of fields, including ...
Body io.ReadCloser
v
... and Body implements the ReadCloser interface ...
type ReadCloser interface {
      Reader
      Closer
}
v
... which implements the Reader interface ...
type Reader interface {
      Read(p []byte) (n int, err error)
}

*/
