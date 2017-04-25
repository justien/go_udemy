// ---- CONSTANTS, IOTA, and BITSHIFTING

package main

import "fmt"

const pie = 3.14
const Language = "golang"

const (
	A = iota
	B
)

const (
	H = iota
	_
	J = iota * 10
)

const (
	_  = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
	TB = 1 << (iota * 10)
)

func main() {

	fmt.Println()

	fmt.Println("----[ Constant, iotas, and bitshifting")

	fmt.Println()

	fmt.Println("Pie has the constant value ", pie)
	fmt.Println("Language has the constant value ", Language)

	fmt.Println()

	fmt.Println("A iota is ", A)
	fmt.Println("B iota is ", B)

	fmt.Println()

	fmt.Println("H iota is ", H)
	fmt.Println("then we have a null variable")
	fmt.Println("J iota is ", J)

	fmt.Println()

	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t\t%d\n", KB, KB)
	fmt.Printf("%b\t%d\n", MB, MB)
	fmt.Printf("%b\t%d\n", GB, GB)
	fmt.Printf("%b\t%d\n", TB, TB)

	fmt.Println()

}

/*
why is this useful???
maybe ... maybe often there are situations were you want to
increment by one ... no wait that doesn't help here bc these
are constants ... urghhh Todd wtf is this for mate????


Ok things that we have here which are new:
----[ Null variable as throw away in iota counts
Use the mull variable to increment the iota w/o wasting a
variable name. Seems like a dodgy shortcut but I guess once
we get in to the idiom, this might make more sense.
....  Or it's some dodgy Todd hack ;)

----[ Bit Shifting
Number followed by << shifts bits to the left, expressed as binary
ex 1 << 4 --> 10000
ex 7 << 4 --> 1110000

*/
