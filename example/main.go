package main

import (
	"fmt"

	v "github.com/bytetwiddler/aba-routingnumber-validator"
)

func main() {

	// ignoring errors for clarity

	// turn string into slice of slice []int required by
	// ValidateAbaRoutingNumber().
	//
	// note that you could just generate a slice []int yourself
	// and pass to ValidateAbaRoutingNumber(). StringSlicer is
	// provided as a convience as folks get stumped by go
	// assuming 012312312 is an out of bound octal.
	// this is a valid routing number
	slice, _ := v.StringSlicer("123123123")

	// ValidateAbaRoutingNumber is the only require function
	// it takes a slice of in []int and returns a bool, error
	res, _ := v.ValidateAbaRoutingNumber(slice)

	fmt.Println(res)
}
