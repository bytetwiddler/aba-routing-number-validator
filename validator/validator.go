package validator

import (
	"fmt"
	"os"
	"strconv"
)

// reversInt reverses and the order of []int slice
func reverseInt(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// StringSlicer takes string of numbers and returns slice []int
func StringSlicer(str string) ([]int, error) {
	if _, err := strconv.Atoi(str); err != nil {
		fmt.Fprintf(os.Stderr, "could not convert '%v' to int\n", str)
		return []int{}, err
	}
	var slice []int // empty slice
	for _, digit := range str {
		slice = append(slice, int(digit)-int('0')) // build up slice
	}
	return slice, nil
}

// slicer takes integer and return []int for each number
func slicer(n int) []int {
	var slc []int
	for n != 0 {
		slc = append(slc, n%10)
		n /= 10
	}

	// now reverse the order of the []slice to put the index order in the order of
	// the submitted integer
	reverseInt(slc)

	return slc
}

// ValidateAbaRoutingNumber uses ABA rules to verify if a routing number is valid
func ValidateAbaRoutingNumber(slc []int) (bool, error) {
	// ABA Routing numbers are alwasy 9 digits
	if len(slc) != 9 {
		return false, fmt.Errorf("routing number slice `%v` is not 9 digits long it is %v digits", slc, len(slc))
	}

	// check of first to digits are between 01-12
	// if first digit is 0 then second digit must be 1,2
	if slc[0] != 0 && slc[0] != 1 {
		return false, fmt.Errorf("routing number slice: %v slice[0]:'%v' is not '0' or '1'", slc, slc[0])
	}
	// if first digit is 1 then second digit must be 1,2
	if slc[0] == 1 {
		if slc[1] != 1 && slc[1] != 2 {
			return false, fmt.Errorf("routing number slice: %v slice[0] is 1 but slice[1]:'%v' is not '1' or '2'", slc, slc[1])
		}
	}

	// multiply the first 8 digits by 7,8,9,7,8,9,7,3 ignore the 9th chksum digit for now
	// while we could do this in a loop we choose to do it explicitly for clarity
	sum := 0
	sum = (slc[0] * 7) + sum
	sum = (slc[1] * 3) + sum
	sum = (slc[2] * 9) + sum
	sum = (slc[3] * 7) + sum
	sum = (slc[4] * 3) + sum
	sum = (slc[5] * 9) + sum
	sum = (slc[6] * 7) + sum
	sum = (slc[7] * 3) + sum

	// turn sum into []int slice so we can get the last digit for the chksum
	chkSlc := slicer(sum)
	// get last digit
	idx := len(chkSlc) - 1
	// compare last digit of routing number to last digit of sum
	return chkSlc[idx] == slc[8], nil
}
