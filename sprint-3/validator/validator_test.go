// Colin McClintic
// Applied Programming Languages
// Coding Standards Validator
// 2/14/2021

package main

import( "testing"
		"fmt"
	 	"codingvalidator.com/val1"
		"codingvalidator.com/val2"
)

func TestValidator(t *testing.T) {
	input := ""
	var longlines []string
	var spacelines []string
	fmt.Println("Executing test00...")
	fmt.Println()
	//test validator
	input = "val.go"
	fmt.Println("Checking val.go for lines over 100 characters...")
	fmt.Println()
	//calls method from val1 module to check standard 1
	longlines = val1.HundredChars(input)
	//case if no lines were detected as longer than 100 characters
	if len(longlines) == 0 {
		fmt.Println("Your code passes this standard. There are 0 lines in your code with over 100 characters")
		fmt.Println()
	}
	//case if lines were detected as longer than 100 characters
	if len(longlines) > 0 {
		fmt.Println("Your code does not pass this standard. The following lines exceed 100 characters:")
		//prints out the number of each line
		for _, line := range longlines {
			fmt.Print(line + " ")
		}
		fmt.Println()
	}
	//display error message if expected outcome is not met
	if len(longlines) != 6 {
   	   	t.Error("Expected output was not received, test failed.")
	}
	if len(longlines) == 6 {
		fmt.Println("Expected output was received, test succeeded.")
	}
	fmt.Println("\nChecking val.go for 3 spaces or more...")
	fmt.Println()
	//calls method from val2 to check standard 2
	spacelines = val2.Spaces(input)
	//case if no lines were detected as having 3 spaces or more
	if len(spacelines) == 0 {
		fmt.Println("Your code passes this standard. There are 0 lines in your code with unnecessary spaces")
		fmt.Println()
	}
	//case if lines were detected as having 3 spaces or more
	if len(spacelines) > 0 {
		fmt.Println("Your code does not pass this standard. The following lines have unnecessary spaces:")
		//prints out the number of each line
		for _, line := range spacelines {
			fmt.Print(line + " ")
		}
		fmt.Println()
	}
	//display error message if expected outcome is not met
   	if len(spacelines) != 0 {
       	t.Error("Expected output was not received, test failed.")
    }
	if len(spacelines) == 0 {
		fmt.Println("Expected output was received, test succeeded.")
	}
	
	fmt.Println("Running automated tests:")
	fmt.Println("Executing test01...")
	fmt.Println()
	//test where both succeed
	input = "testcode01.go"
	fmt.Println("Checking your code for lines over 100 characters...")
	fmt.Println()
	//calls method from val1 module to check standard 1
	longlines = val1.HundredChars(input)
	//case if no lines were detected as longer than 100 characters
	if len(longlines) == 0 {
		fmt.Println("Your code passes this standard. There are 0 lines in your code with over 100 characters")
		fmt.Println()
	}
	//case if lines were detected as longer than 100 characters
	if len(longlines) > 0 {
		fmt.Println("Your code does not pass this standard. The following lines exceed 100 characters:")
		//prints out the number of each line
		for _, line := range longlines {
			fmt.Print(line + " ")
		}
		fmt.Println()
	}
	//display error message if expected outcome is not met
	if len(longlines) != 0 {
        t.Error("Expected output was not received, test failed.")
    }
	if len(longlines) == 0 {
		fmt.Println("Expected output was received, test succeeded.")
	}
	fmt.Println("\nChecking your code for 3 spaces or more...")
	fmt.Println()
	//calls method from val2 to check standard 2
	spacelines = val2.Spaces(input)
	//case if no lines were detected as having 3 spaces or more
	if len(spacelines) == 0 {
		fmt.Println("Your code passes this standard. There are 0 lines in your code with unnecessary spaces")
		fmt.Println()
	}
	//case if lines were detected as having 3 spaces or more
	if len(spacelines) > 0 {
		fmt.Println("Your code does not pass this standard. The following lines have unnecessary spaces:")
		//prints out the number of each line
		for _, line := range spacelines {
			fmt.Print(line + " ")
		}
		fmt.Println()
	}
	//display error message if expected outcome is not met
    if len(spacelines) != 0 {
        t.Error("Expected output was not received, test failed.")
    }
	if len(spacelines) == 0 {
		fmt.Println("Expected output was received, test succeeded.")
	}
	fmt.Println("\nExecuting test02...")
	//test both failing
	input = "testcode02.go"
	fmt.Println("\nChecking your code for lines over 100 characters...")
	fmt.Println()
	//calls method from val1 module to check standard 1
	longlines = val1.HundredChars(input)
	//case if no lines were detected as longer than 100 characters
	if len(longlines) == 0 {
		fmt.Println("Your code passes this standard. There are 0 lines in your code with over 100 characters")
		fmt.Println()
	}
	//case if lines were detected as longer than 100 characters
	if len(longlines) > 0 {
		fmt.Println("Your code does not pass this standard. The following lines exceed 100 characters:")
		//prints out the number of each line
		for _, line := range longlines {
			fmt.Print(line + " ")
		}
		fmt.Println()
	}
	//display error message if expected outcome is not met
	if len(longlines) != 1 {
        t.Error("Expected output was not received, test failed.")
    }
	if len(longlines) == 1 {
		fmt.Println("Expected output was received, test succeeded.")
	}
	fmt.Println("\nChecking your code for 3 spaces or more...")
	fmt.Println()
	//calls method from val2 to check standard 2
	spacelines = val2.Spaces(input)
	//case if no lines were detected as having 3 spaces or more
	if len(spacelines) == 0 {
		fmt.Println("Your code passes this standard. There are 0 lines in your code with unnecessary spaces")
		fmt.Println()
	}
	//case if lines were detected as having 3 spaces or more
	if len(spacelines) > 0 {
		fmt.Println("Your code does not pass this standard. The following lines have unnecessary spaces:")
		//prints out the number of each line
		for _, line := range spacelines {
			fmt.Print(line + " ")
		}
		fmt.Println()
	}
    if len(spacelines) != 1 {
        t.Error("Expected output was not received, test failed.")
    }
	//display error message if expected outcome is not met
	if len(spacelines) == 1 {
		fmt.Println("Expected output was received, test succeeded.")
	}
	fmt.Println("\nExecuting test03...")
	//test val1 failing
	input = "testcode03.go"
	fmt.Println("Checking your code for lines over 100 characters...")
	fmt.Println()
	//calls method from val1 module to check standard 1
	longlines = val1.HundredChars(input)
	//case if no lines were detected as longer than 100 characters
	if len(longlines) == 0 {
		fmt.Println("Your code passes this standard. There are 0 lines in your code with over 100 characters")
		fmt.Println()
	}
	//case if lines were detected as longer than 100 characters
	if len(longlines) > 0 {
		fmt.Println("Your code does not pass this standard. The following lines exceed 100 characters:")
		//prints out the number of each line
		for _, line := range longlines {
			fmt.Print(line + " ")
		}
		fmt.Println()
	}
	//display error message if expected outcome is not met
	if len(longlines) != 1 {
        t.Error("Expected output was not received, test failed.")
    }
	if len(longlines) == 1 {
		fmt.Println("Expected output was received, test succeeded.")
	}
	fmt.Println("\nChecking your code for 3 spaces or more...")
	fmt.Println()
	//calls method from val2 to check standard 2
	spacelines = val2.Spaces(input)
	//case if no lines were detected as having 3 spaces or more
	if len(spacelines) == 0 {
		fmt.Println("Your code passes this standard. There are 0 lines in your code with unnecessary spaces")
		fmt.Println()
	}
	//case if lines were detected as having 3 spaces or more
	if len(spacelines) > 0 {
		fmt.Println("Your code does not pass this standard. The following lines have unnecessary spaces:")
		//prints out the number of each line
		for _, line := range spacelines {
			fmt.Print(line + " ")
		}
		fmt.Println()
	}
	//display error message if expected outcome is not met
    if len(spacelines) != 0 {
        t.Error("Expected output was not received, test failed.")
    }
	if len(spacelines) == 0 {
		fmt.Println("Expected output was received, test succeeded.")
	}
	fmt.Println("\nExecuting test04...")
	//test val2 failing
	input = "testcode04.go"
	fmt.Println("\nChecking your code for lines over 100 characters...")
	fmt.Println()
	//calls method from val1 module to check standard 1
	longlines = val1.HundredChars(input)
	//case if no lines were detected as longer than 100 characters
	if len(longlines) == 0 {
		fmt.Println("Your code passes this standard. There are 0 lines in your code with over 100 characters")
		fmt.Println()
	}
	//case if lines were detected as longer than 100 characters
	if len(longlines) > 0 {
		fmt.Println("Your code does not pass this standard. The following lines exceed 100 characters:")
		//prints out the number of each line
		for _, line := range longlines {
			fmt.Print(line + " ")
		}
		fmt.Println()
	}
	//display error message if expected outcome is not met
	if len(longlines) != 0 {
        t.Error("Expected output was not received, test failed.")
    }
	if len(longlines) == 0 {
		fmt.Println("Expected output was received, test succeeded.")
	}
	fmt.Println("\nChecking your code for 3 spaces or more...")
	fmt.Println()
	//calls method from val2 to check standard 2
	spacelines = val2.Spaces(input)
	//case if no lines were detected as having 3 spaces or more
	if len(spacelines) == 0 {
		fmt.Println("Your code passes this standard. There are 0 lines in your code with unnecessary spaces")
		fmt.Println()
	}
	//case if lines were detected as having 3 spaces or more
	if len(spacelines) > 0 {
		fmt.Println("Your code does not pass this standard. The following lines have unnecessary spaces:")
		//prints out the number of each line
		for _, line := range spacelines {
			fmt.Print(line + " ")
		}
		fmt.Println()
	}
    if len(spacelines) != 1 {
        t.Error("Expected output was not received, test failed.")
    }
	if len(spacelines) == 1 {
		fmt.Println("Expected output was received, test succeeded.")
	}
}
