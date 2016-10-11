// Reads in an HCL file and converts it to JSON on stdout or the specified file.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/hashicorp/hcl"
)

func main() {
	// Check arguments
	if len(os.Args) <= 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s <input file> [output file]\n", path.Base(os.Args[0]))
		os.Exit(1)
	}

	// Read in input file
	input, err := ioutil.ReadFile(os.Args[1])
	errorOut(err, 2)

	// Parse HCL
	var jsonStruct interface{}
	errorOut(hcl.Unmarshal(input, &jsonStruct), 3)

	// Convert parsed data to JSON
	jsonText, err := json.MarshalIndent(jsonStruct, "", "    ")
	errorOut(err, 4)

	// Decide where to send the converted file
	output := os.Stdout
	if len(os.Args) == 3 {
		output, err = os.OpenFile(os.Args[2], os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		errorOut(err, 5)
		defer output.Close()
	}

	fmt.Fprintln(output, string(jsonText))
}

func errorOut(err error, exitCode int) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(exitCode)
	}
}
