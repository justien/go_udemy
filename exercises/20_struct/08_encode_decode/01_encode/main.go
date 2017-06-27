package main

import (
	"encoding/json"
	"os"
)

type person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {

	p1 := person{"James", "Bond", 20, 007}

	//Hey Jason, plz NewEncode to standard out the stuff that
	//you get when you Encode p1 thx!
	json.NewEncoder(os.Stdout).Encode(p1)
}

// encoding and decoding passes things to readers and writers!
