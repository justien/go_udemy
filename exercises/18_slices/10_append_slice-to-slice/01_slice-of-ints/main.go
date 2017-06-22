package main

import "fmt"

func main() {

	mySlice00 := []int{1, 2, 3, 5, 18}
	mySlice01 := []int{12, 13, 14, 15}

	mySlice00 = append(mySlice00, mySlice01...)

	fmt.Println(mySlice00)
}
