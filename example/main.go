package main

import (
	"fmt"

	v "github.com/bytetwiddler/aba-routingnumber-validator"
)

func main() {

	// ignoring errors for clarity

	// note: that you could just generate a slice of []int 
	// and pass to ValidateAbaRoutingNumber(). 
	// StringSlicer function is provided as a convience as
	// folks get stumped by golang assuming 012312312 is an
	// out of bound octal.
	// 123123123 is a valid routing number.
	
	// ValidateAbaRoutingNumber()
	// turns string version of the number into slice of []int
	slice, _ := v.StringSlicer("123123123")

	// ValidateAbaRoutingNumber is the only required function
	// it takes a slice of int []int and returns a bool, error
	res, _ := v.ValidateAbaRoutingNumber(slice)

	fmt.Println(res)
}
