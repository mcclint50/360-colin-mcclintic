// Colin McClintic
// Applied Programming Languages
// Coding Standards Validator
// 2/14/2021

package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"codingvalidator.com/val1"		//imports module for testing standard 1
	"codingvalidator.com/val2"		//imports module for testing standard 2
)

func main() {
	//declare reader and scanner
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)
	restart := "y"
	//main program loop
	for restart == "y" {
		standard := "val help"
		fmt.Println("Welcome to the Coding Standards Validator!\n")
		fmt.Println("1. 100 characters per line")
		fmt.Println("2. Tabs used for indenting\n")
		//loops until user has selected a standard
		for standard == "val help" {
			//read standard input from user
			fmt.Println("Enter the number of the coding standard you would like validating or enter 'val help' for further instructions: ")
			standard, _ = reader.ReadString('\n')
			standard = strings.Replace(standard, "\r\n", "", -1)
			//prints val help message
			if (standard == "val help") {
				fmt.Println("\nThis program is designed for you to automatically test your code. Input the number that corresponds to the coding")
				fmt.Println("standard you would like to verify. The first one checks if each line in your code is wrapped at 100 characters.")
				fmt.Println("The second one checks if your code uses tabs to indent instead of spaces.\n")
			}
		}
		//get source code location from user
		fmt.Println("Enter the path of the source code file you want to validate: ")
		scanner.Scan()
		input := scanner.Text()
		//standard 1
		if (standard == "1") {
			fmt.Println("\nChecking your code for lines over 100 characters...\n")
			//calls method from val1 module to check standard 1
			longlines := val1.HundredChars(input)
			//case if no lines were detected as longer than 100 characters
			if len(longlines) == 0 {
				fmt.Println("Your code passes this standard. There are 0 lines in your code with over 100 characters")
			}
			//case if lines were detected as longer than 100 characters
			if len(longlines) > 0 {
				fmt.Println("Your code does not pass this standard. The following lines exceed 100 characters:")
				//prints out the number of each line
				for _, line := range longlines {
					fmt.Print(line + " ")
				}
				fmt.Println("")
			}
		}
		//standard 2
		if (standard == "2") {
			fmt.Println("\nChecking your code for 3 spaces or more...\n")
			//calls method from val2 to check standard 2
			spacelines := val2.Spaces(input)
			//case if no lines were detected as having 3 spaces or more
			if len(spacelines) == 0 {
				fmt.Println("Your code passes this standard. There are 0 lines in your code with unnecessary spaces")
			}
			//case if lines were detected as having 3 spaces or more
			if len(spacelines) > 0 {
				fmt.Println("Your code does not pass this standard. The following lines have unnecessary spaces:")
				//prints out the number of each line
				for _, line := range spacelines {
					fmt.Print(line + " ")
				}
				fmt.Println("")
			}
		}
		//ask user if they would like to restart the program
		fmt.Print("\nWould you like to restart the program (y/n)? ")
		scanner.Scan()
		restart = scanner.Text()
		fmt.Println("")
	}
}

