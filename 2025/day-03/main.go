package aoc

import (
	"fmt"
	"strconv"
)

/*
	cd 2025/day-03/
	go run ./cmd/
*/

func Run() {
	fmt.Println("day03")
}

func GetIndexAndValueOfBiggestNumberFromRange(input string) (int, string) {

	returnIndex, returnValue := -1, -1

	for index, r := range input {
		character := fmt.Sprintf("%c", r)
		value, _ := strconv.Atoi(character)
		if value > returnValue {
			returnValue = value
			returnIndex = index			
		}
	}

	return returnIndex, strconv.Itoa(returnValue)
}

func GetBiggestNumberFromRange(input string) string {

	firstBiggestIndex, firstBiggestValue := GetIndexAndValueOfBiggestNumberFromRange(input)
	lengthOfInput := len(input)
	subset := ""

	if firstBiggestIndex == (lengthOfInput - 1) {
		// biggest character was in last place
		subset = input[:firstBiggestIndex]
	} else {
		subset = input[firstBiggestIndex+1:]
	}

	_, secondBiggestValue := GetIndexAndValueOfBiggestNumberFromRange(subset)

	returnString := ""	
	 if firstBiggestIndex == (lengthOfInput - 1) { 
		returnString = fmt.Sprintf("%s%s", secondBiggestValue, firstBiggestValue)
	 } else {
		returnString = fmt.Sprintf("%s%s", firstBiggestValue, secondBiggestValue)
	 }

	return returnString
}

