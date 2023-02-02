package main

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/table"
)

type test struct {
	value       string
	expect      bool
	description string
}

func tablePrint(tw table.Writer) {
	// append a header row
	tw.AppendHeader(table.Row{"PASS/FAIL:\u2713/\u2717", "Input", "Expected", "Received", "Failure description"})
	// render table
	fmt.Printf("Unit Results:\n%s\n", tw.Render())
}

/*
func reverseInt(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// Since Go assumes and number starting with 0 is octal
// we pass in a string
func stringSlicer(str string) []int {
	if _, err := strconv.Atoi(str); err != nil {
		fmt.Fprintf(os.Stderr, "could not convert '%v' to int\n", str)
	}
	var slice []int // empty slice
	for _, digit := range str {
		slice = append(slice, int(digit)-int('0')) // build up slice
	}
	return slice
}

// this was my original but ran into issues with octals when testing
// with a routing number starting with 0
func slicer(n int) []int {
	var slc []int
	for n != 0 {
		slc = append(slc, n%10)
		n /= 10
	}
	reverseInt(slc)
	return slc
}

func validateAbaRoutingNumber(slc []int) bool {

	// ABA Routing numbers are alwasy 9 digits
	if len(slc) != 9 {
		if verbose {
			fmt.Printf("routing number slice `%v` is not 9 digits, it is %v digits \n", slc, len(slc))
		}
		return false
	}

	// check of first to digits are between 01-12
	// if first digit is 0 then second digit must be 1,2
	if slc[0] != 0 && slc[0] != 1 {
		if verbose {
			fmt.Printf("routing number slice: %v slice[0]:'%v' is not '0' or '1'\n", slc, slc[0])
		}
		return false
	}
	// if first digit is 1 then second digit must be 1,2
	if slc[0] == 1 {
		if slc[1] != 1 && slc[1] != 2 {
			if verbose {
				fmt.Printf("routing number slice: %v slice[0] is 1 but slice[1]:'%v' is not '1' or '2'\n", slc, slc[1])
			}
			return false
		}
	}

	// multiply the first 8 digits by 7,8,9,7,8,9,7,3 ignore the 9th chksum digit for now
	// do it explicity for clarity
	sum := 0
	sum = (slc[0] * 7) + sum
	sum = (slc[1] * 3) + sum
	sum = (slc[2] * 9) + sum
	sum = (slc[3] * 7) + sum
	sum = (slc[4] * 3) + sum
	sum = (slc[5] * 9) + sum
	sum = (slc[6] * 7) + sum
	sum = (slc[7] * 3) + sum

	// turn sum into slice so we can get the last digit
	chkSlc := slicer(sum)

	// get last digit
	idx := len(chkSlc) - 1

	// compare last digit of routing number to last digit of sum
	return chkSlc[idx] == slc[8]
}
*/
var verbose = false

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "-v" {
			verbose = true
		}
	}

	// setup our test table
	testTable := []test{
		{"123123123", true, ""},
		{"122105278", true, ""},
		{"022105278", false, "check sum fail"},
		{"12345678", false, "short 8 digits"},
		{"1234567890", false, "long 10 digits"},
		{"934567890", false, "1st 2 digits not between 01-12"},
		{"134567890", false, "1st 2 digits not between 01-12"},
	}

	tw := table.NewWriter()
	for _, t := range testTable {
		slc := stringSlicer(t.value)
		result := validateAbaRoutingNumber(slc)
		if result == t.expect {
			tw.AppendRow(table.Row{"\u2713", t.value, t.expect, result, t.description})
		} else {
			tw.AppendRow(table.Row{"\u2717", t.value, t.expect, result, t.description})
		}
	}

	tablePrint(tw)
}
