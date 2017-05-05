package main

import "fmt"

func main() {
	fmt.Println([]byte("Hello"))
	for i := 5000; i <= 5010; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
	for i := 18765; i <= 18775; i++ {
		fmt.Printf("%v - %v - %v\n", i, string(i), []byte(string(i)))
	}
}
