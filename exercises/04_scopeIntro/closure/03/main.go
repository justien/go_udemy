// package main
//
// import "fmt"
//
// func wrapper() func() int {
// 	x := 0
// 	return func() int { // <-- an anonymous function
// 		x++
// 		return x
// 	}
// }
//
// func main() {
// 	increment := wrapper()
// 	fmt.Println("Water puddled under old cars ", increment())
// 	fmt.Println("Water puddled under abandoned cars ", increment())
// 	fmt.Println("Water puddled under rusted cars ", wrapper())
// 	fmt.Println("The value of x is unknown to func main")
//
// 	soaring := wrapper()
// 	fmt.Println("The air was filled with dust ", soaring())
// 	fmt.Println("The air was filled with a low sound ", soaring())
// }

/*
|    In this script, x is part of a function called wrapper.
|    wrapper is a small func that:
|      -- declares and initialises x
|      -- returns a func
... where is this function that we're returning???
  --> it's found after the word 'return'
	    .... it is the anonymous function wee-oo-wee-oo

	   -- so when wrapper gets called,
		 func() int {
		    x++
		    return x
	    }
		    .... is what gets returned.

... and this function is what is assigned to increment!

I don't fully get this yet but I will later on!

|
*/
