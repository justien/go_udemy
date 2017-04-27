package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {

	fmt.Println()
	fmt.Println("Section 7: Converting meters to yards")
	fmt.Println()

	var meters float64
	fmt.Print("Enter meters that you swam: ")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, " meters is equivalent to ", yards, " yards.")
	fmt.Println()
	fmt.Println()
}
