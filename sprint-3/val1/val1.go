// Colin McClintic
// Applied Programming Languages
// Coding Standards Validator
// 2/14/2021

package val1

import (
	"bufio"
	"log"
	"os"
	"fmt"
)

//function that checks each line of code for more than 100 chars
func HundredChars(code string) ([]string) {
	//open file
	file, err := os.Open(code)
 
	//handles error opening file
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
 
	//scans each line of the file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	//initialize lists of strings that we'll need and count
	var txtlines []string
	var toolong []string
	count := 1

	//adds each line to list of txtlines
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	//closes file
	file.Close()
 
	//checks if each line is over 100 chars
	for _, eachline := range txtlines {
		if (len(eachline) > 100) {
			//adds number of line to list of lines that are too long
			toolong = append(toolong, fmt.Sprint(count))
		}
		count = count + 1
	}
	//returns list of numbers of lines that are too long
	return toolong
}