package main

import "fmt"

func main() {

	mySlice := make([]int, 0, 3)

	fmt.Println("---------------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("--------------------")

	for i := 0; i < 15; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Length:", len(mySlice), "Cap:", cap(mySlice), "Value:", mySlice[i])
	}
}
