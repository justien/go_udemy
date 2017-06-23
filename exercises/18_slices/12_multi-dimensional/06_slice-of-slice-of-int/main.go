package main

import "fmt"

func main() {

	// Create a 2d slice and populate it with a loop

	// Create 2d slice with 1° index length 0 and cap 3
	columns := make([][]int, 0, 3)
	// row := make([]int, 0, 4) DON'T PUT ROW HERE

	// Nested loop to populate 2d slice
	for i := 0; i < 3; i++ {

		row := make([]int, 0, 3) // gimme a clean row!

		// this loop adds the value of j to the row slice
		for j := 7; j < 10; j++ {
			row = append(row, j)
			fmt.Println("current row entry:", row[(j-7)])
			fmt.Println("current state of row:", row)
		}

		columns = append(columns, row)
		fmt.Println("The current columns:", columns)
		fmt.Println()
	}

	fmt.Println("\n\nThe final slice:", columns)

}
