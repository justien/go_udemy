package main

import "fmt"

func main() {

	m := make([]string, 2, 25)

	fmt.Println("This is what the m slice looks like atm:", m)

	changeMe(m)

	fmt.Println("Now slice m looks like this:", m)

}

func changeMe(z []string) {
	z[0] = "Todd"
	fmt.Println(z)
}
