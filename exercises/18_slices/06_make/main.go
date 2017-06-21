package main

import "fmt"

func main() {

	custNo := make([]int, 3)

	custNo[0] = 7
	custNo[1] = 10
	custNo[2] = 15

	fmt.Println(custNo[0])
	fmt.Println(custNo[1])
	fmt.Println(custNo[2])

	greeting := make([]string, 3, 5)

	greeting[0] = "good Morning"
	greeting[1] = "Guten Tag!"
	greeting[2] = "Selamat datang!"

	fmt.Println(greeting[2])
}
