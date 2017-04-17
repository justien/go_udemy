// Package stringutil contains utility functions for working with strings.
package stringutil

//Reverse returns its arfumet string reversed rune-wise left to right.
func Reverse(s string) string {
	return reverseTwo(s)
}

/*
go build
  go build reverse.go reverseTwo.go
  won't produce an output file.

go install
  will place the packeage inside the pkg directory of the workspace.
*/
