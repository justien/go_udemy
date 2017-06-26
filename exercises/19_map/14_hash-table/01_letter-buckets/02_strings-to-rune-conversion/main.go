package main

import "fmt"

func main() {
	letter := rune("Hello"[0]) //strings are a slice of bytes
	fmt.Println(letter)
}

/*
strings are a slice of bytes,
so this will output the byte value of the first
letter of the string
*/
