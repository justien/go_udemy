package main

import "fmt"

func main() {

	halvies := func(z int) {
		divide := z / 2
		remainder := z % 2
		even := true
		if remainder == 1 {
			even = false
		}
		fmt.Println(divide, even)
	}

	halvies(7)
}

// func half(z int) {
//
// 	divide := z / 2
// 	remainder := z & 2
// 	even := true
// 	if remainder == 1 {
// 		even = false
// 	}
// 	fmt.Println(divide, even)
// }
