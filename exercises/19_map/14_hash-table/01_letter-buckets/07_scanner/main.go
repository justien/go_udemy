package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	const bunch_o_words = "It is not the critic who counts; not the man who points out how the strong man stumbles, or where the doer of deeds could have done them better. The credit belongs to the man who is actually in the arena, whose face is marred by dust and sweat and blood; who strives valiantly; who errs, who comes short again and again, because there is no effort without error and shortcoming; but who does actually strive to do the deeds; who knows great enthusiasms, the great devotions; who spends himself in a worthy cause; who at the best knows in the end the triumph of high achievement, and who at the worst, if he fails, at least fails while daring greatly, so that his place shall never be with those cold and timid souls who neither know victory nor defeat. "

	/*
			I don't get how this works :(
		  ... apparently NewScanner returns a pointer to a scanner
		  and then when I have a pointer to a scanner, I can do Split and
		  pass in to Split bufio.ScanWords
	*/

	scanner := bufio.NewScanner(strings.NewReader(bunch_o_words))

	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)

	/*
		Count the words.
		Scan() returns a bool, so as long as there is more stuff
		in scanner, the for loop will be true, and will kepp running.
	  Scan() returns false when it reaches the end of a file.
	*/

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading bunch_o_words:", err)
	}
}
