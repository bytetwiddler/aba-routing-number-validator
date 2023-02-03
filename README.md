# aba-routing-number-validator
Validate ABA bank routing numbers

We took the requiments from the following videos.

https://www.youtube.com/watch?v=5wAE3E5zBe4

https://www.youtube.com/watch?v=D0jY_CuzcGw

It uses the 73973972 weights + check digit algorithm

Little to no testing was done beyond the unit test 
provided. 

**example**
```package main

import (
	"fmt"

	v "github.com/bytetwiddler/aba-routingnumber-validator"
)


/* 
 * Ignoring errors for clarity.
 * One could just generate a slice of int and pass to ValidateAbaRoutingNumber() directly.
 * The StringSlicer() function is provided as a convience as folks tend to get stumped by 
 * golang assuming 012312312 is an out of bound octal.
 */ 
 
func main() {

	// turns string version of the number into slice of []int
	slice, _ := v.StringSlicer("123123123") // this is valid 

	// ValidateAbaRoutingNumber is the only required function
	// it takes a slice of int []int and returns a bool, error
	res, _ := v.ValidateAbaRoutingNumber(slice)

	fmt.Println(res)
}
```
