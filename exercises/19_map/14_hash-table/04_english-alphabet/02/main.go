package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	// give me the response and error when I call http.Get(randoWebSite)

	response, error := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if error != nil {
		log.Fatalln(error)
	}

	// make a map to hold all the words retrieved

	wordsMap := make(map[string]string)

	/*
	 Make me an object which is of type
	 "buffered fresh scan" of the response from the web.

	 Now you have a big blob of web stuff.

	 Using the Split method, chop it up by words.
	 I don't understand why bufio has to be there.

	*/

	myWordList := bufio.NewScanner(response.Body)
	myWordList.Split(bufio.ScanWords) // take the list and split it by words

	/*
	   This for loop scans through the list of words.
	   It takes each element of text found, and puts it into a map.
	   The key will be the word, and the value will be nil.
	*/

	for myWordList.Scan() {
		wordsMap[myWordList.Text()] = ""
	}

	/*
			----
			myWordList also contains an entry relating to errors from the http.get()
		  If this entry is not nil, tell the user about it.
	*/

	if error := myWordList.Err(); error != nil {
		fmt.Fprintln(os.Stderr, "reading input:", error)
	}

	/*
		----
			Nested loop that counts to 200.
			Whilist it does that, it also ranges over
			the map of words, and prints out the entry
	*/

	i := 0
	for k := range wordsMap {
		fmt.Println(k)
		if i == 18 {
			break
		}
		i++
	}

}
