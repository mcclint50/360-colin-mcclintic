// Colin McClintic
// Applied Programming Languages
// Match Maker
// 1/31/2021
package main

import (
	"fmt"
	"math"
)

const myval1, myval2, myval3, myval4, myval5 int = 5, 1, 4, 4, 5
const weight1, weight2, weight3, weight4, weight5 int = 2, 1, 3, 2, 3
const trueLove, justFriends, runAway int = 100, 90, 80

//Validate function
func validate(val int) bool {
	if (val < 1 || val > 6) {
		return true
	}
	if (val > 0 || val < 7) {
		return false
	}
}

func main() {
	//Initialize values
	var again = "y"
	var val1, val2, val3, val4, val5 int = 0, 0, 0, 0, 0

	//Header
	fmt.Println("**************************************************************************************")
	fmt.Println("                                Match Maker Application")
	fmt.Println("                              Created by: Colin McClintic")
	fmt.Println("**************************************************************************************")
	fmt.Println("")

	//Play again loop
	for again == "y" {
		fmt.Println("Welcome to the Match Maker application. This program will tell you how well")
		fmt.Println("of a match you and I are. You will answer 5 questions. To each question,")
		fmt.Println("answer 5 if you strongly agree, 4 if you agree, 3 if you neither agree")
		fmt.Println("nor disagree, 2 if you disagree, and 1 if you strongly disagree.")
		fmt.Println("")
		fmt.Println("Here we go...")
		fmt.Println("")

		//Question 1
		//Asks again if answer is out of range
		for (validate(val1)) {
			fmt.Println("Russell Westbrook is the best point guard in the league.")
			fmt.Scanln(&val1)
			//Displays error message if value is out of range
			if (validate(val1)) {
				fmt.Println("Value entered is out of range. Please enter a value between 1 and 5")
				fmt.Println("")
			}
		}
		fmt.Println("")

		//Question 2
		for validate(val2) {
			fmt.Println("Asparagus is icky.")
			fmt.Scanln(&val2)
			if validate(val2) {
				fmt.Println("Value entered is out of range. Please enter a value between 1 and 5")
				fmt.Println("")
			}
		}
		fmt.Println("")

		//Question 3
		for validate(val3) {
			fmt.Println("If you start something, you must finish it.")
			fmt.Scanln(&val3)
			if validate(val3) {
				fmt.Println("Value entered is out of range. Please enter a value between 1 and 5")
				fmt.Println("")
			}
		}
		fmt.Println("")

		//Question 4
		for validate(val4) {
			fmt.Println("My significant other must have the same music taste as me.")
			fmt.Scanln(&val4)
			if validate(val4) {
				fmt.Println("Value entered is out of range. Please enter a value between 1 and 5")
				fmt.Println("")
			}
		}
		fmt.Println("")

		//Question 5
		for validate(val5) {
			fmt.Println("I enjoy trying new things.")
			fmt.Scanln(&val5)
			if validate(val5) {
				fmt.Println("Value entered is out of range. Please enter a value between 1 and 5")
				fmt.Println("")
			}
		}

		//Calculates difference between answers and my answers
		var dif1 float64 = float64(val1 - myval1)
		var dif2 float64 = float64(val2 - myval2)
		var dif3 float64 = float64(val3 - myval3)
		var dif4 float64 = float64(val4 - myval4)
		var dif5 float64 = float64(val5 - myval5)

		//Calculates compatibility score
		var score = 100 - ((math.Abs(dif1)*float64(weight1)) + (math.Abs(dif2)*float64(weight2)) + (math.Abs(dif3)*float64(weight3)) + (math.Abs(dif4)*float64(weight4)) + (math.Abs(dif5)*float64(weight5)))
		fmt.Println("")

		//Summary
		fmt.Println("---------------------------------------------------------------------------------------------------------------------")
		fmt.Println("")
		fmt.Println("Summary:     Compatibility Score (0-best, 4-worst)      Question's Weight (1-least important, 3-most important)")
		fmt.Println("Question 1 -             ", math.Abs(dif1), "                                                 ", weight1)
		fmt.Println("Question 2 -             ", math.Abs(dif2), "                                                 ", weight2)
		fmt.Println("Question 3 -             ", math.Abs(dif3), "                                                 ", weight3)
		fmt.Println("Question 4 -             ", math.Abs(dif4), "                                                 ", weight4)
		fmt.Println("Question 5 -             ", math.Abs(dif5), "                                                 ", weight5)
		fmt.Println("")

		//Displays compatibility
		fmt.Println("Our compatiblility is: ", score, "%")

		//Displays closing message based on user's score
		switch {
		case score < float64(runAway):
			fmt.Println("You and I are not compatible at all. It just wasn't meant to be...")
		case score < float64(justFriends):
			fmt.Println("You and I are not soul mates. That doesn't mean we can't still be friends though.")
		case score < float64(trueLove):
			fmt.Println("We have a lot in common! Looks like it's a perfect match <3")
		}
		fmt.Println("")

		//Asks user if they want to restart the program
		fmt.Print("Would you like to restart the program? (y/n) ")
		fmt.Scanln(&again)
	}
}