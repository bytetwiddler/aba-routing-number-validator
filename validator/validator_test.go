package aba-routing-number-validator

import (
	"fmt"
	"testing"

	"github.com/jedib0t/go-pretty/table"
)

type test struct {
	input       string
	want        bool
	description string
}

func tablePrint(tw table.Writer) {
	tw.AppendHeader(table.Row{"PASS/FAIL:\u2713/\u2717", "Input", "Expected", "Received", "Failure description"})
	fmt.Printf("Unit Results:\n%s\n", tw.Render())
}

func TestValidator(t *testing.T) {

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
	for i, tt := range testTable {
		slc, _ := StringSlicer(tt.input)
		result, _ := ValidateAbaRoutingNumber(slc)
		if result == tt.want {
			tw.AppendRow(table.Row{"\u2713", tt.input, tt.want, result, tt.description})
		} else {
			tw.AppendRow(table.Row{"\u2717", tt.input, tt.want, result, tt.description})
			t.Errorf("test %d - %v:  expected: %v got: %v :%v", i+1, tt.input, tt.want, result, tt.description)
		}
	}

	tablePrint(tw)
}
