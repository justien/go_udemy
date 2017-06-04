package main

import "fmt"

func main() {
	data := []float64{13, 23, 29, 31, 2.0}
	avg := average(data)
	fmt.Println("this is the average of the values", avg)
}

func average(sf []float64) float64 {
	var total float64
	for _, v := range sf {
		total += v
		fmt.Println("Current value is", v, "and current total is ", total)
	}
	fmt.Println("The number of values is", len(sf))
	return total / float64(len(sf))
}
