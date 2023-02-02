package validator

import (
	"fmt"
	"os"

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
		result := v.ValidateAbaRoutingNumber(slc)
		if result == t.expect {
			tw.AppendRow(table.Row{"\u2713", t.value, t.expect, result, t.description})
		} else {
			tw.AppendRow(table.Row{"\u2717", t.value, t.expect, result, t.description})
		}
	}

	tablePrint(tw)
}
