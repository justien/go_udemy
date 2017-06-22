package main

import "fmt"

func main() {

	mySlice00 := []string{"Monday", "Tuesday"}
	mySlice01 := []string{"Wednesday", "Thursday", "Friday"}

	mySlice00 = append(mySlice00, mySlice01...)

	fmt.Println(mySlice00)
}
