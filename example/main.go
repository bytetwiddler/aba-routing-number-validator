package main

import (
	"fmt"

	v "github.com/bytetwiddler/aba-routingnumber-validator"
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

func main() {

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

	fmt.Println(`Note: errors are just printed out and not handled as we are passing test data
and expect errors. Check the tables PASS/FAIL field to see if a particular 
routing number got the expected result.
	 `)
	tw := table.NewWriter()
	for _, t := range testTable {
		slc, err := v.StringSlicer(t.value)
		if err != nil {
			fmt.Println("error: ", err)
		}
		result, err := v.ValidateAbaRoutingNumber(slc)
		if err != nil {
			fmt.Println("error: ", err)
		}
		if result == t.expect {
			tw.AppendRow(table.Row{"\u2713", t.value, t.expect, result, t.description})
		} else {
			tw.AppendRow(table.Row{"\u2717", t.value, t.expect, result, t.description})
		}
	}

	tablePrint(tw)
}
