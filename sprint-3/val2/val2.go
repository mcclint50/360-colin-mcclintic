// Colin McClintic
// Applied Programming Languages
// Coding Standards Validator
// 2/14/2021

package val2

import (
	"log"
	"os"
	"bufio"
	"strings"
	"fmt"
)

//function that checks if there are 3 spaces or more in the code
func Spaces(code string) ([] string) {
	//opens file
	file, err := os.Open(code)
 
	//handles error with opening file
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
 
	//scans each line of code
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	//intializes lists of strings that we'll need and count
	var txtlines []string
	var spaces []string
	count := 1

	//adds each line to list of txtlines
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	//closes file
	file.Close()
 
	//checks if line has 3 or more spaces
	for _, eachline := range txtlines {
		if strings.Contains(eachline, "   ") {
			//adds number of line to list of lines that violate the standard
			spaces = append(spaces, fmt.Sprint(count))
		}
		count = count + 1
	}
	//returns list of number of lines that have 3 or more spaces
	return spaces
}