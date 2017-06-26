package main

import "fmt"

/*
count the string value of the ascii character set
and the associated bucket index (ie index value mod 12)
*/

func main() {
	for i := 65; i <= 122; i++ {
		fmt.Println(i, " - ", string(i), " - ", i%12)
	}
}
