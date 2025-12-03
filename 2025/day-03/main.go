package aoc

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
	cd 2025/day-03/
	go run ./cmd/
*/

func Run() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()
		returnedValue := GetBiggestNumberFromRange(currentLine)
		currentNumber, _ := strconv.Atoi(returnedValue)
		sum += currentNumber
	}
	fmt.Println(sum)
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

