package main

import "fmt"

func main() {

	mySlice := []int{1, 2, 3, 5, 8}
	fmt.Printf("%T is the type of mySlice\n", mySlice)
	fmt.Println("And here is mySlice:", mySlice)

	fmt.Println("Here is a slice of mySlice from 2 to 4:", mySlice[2:4])
	fmt.Println("Here is element at index 2 of mySlice:", mySlice[2])

	mySlice01 := []string{"a", "b", "c", "d", "e"}
	fmt.Printf("%T is the type of mySlice01\n", mySlice01)
	fmt.Println("And here is mySlice01:", mySlice01)

	fmt.Println("Here is a slice of mySlice01 from 2 to 4:", mySlice01[2:4])
	fmt.Println("Here is element at index 2 of mySlice01:", mySlice01[2])
	fmt.Println("On the next line, \"mySlice01\"[2] gives us the byte value of S")
	fmt.Println("mySlice01"[2])

	fmt.Println("")
}
