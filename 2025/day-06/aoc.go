package aoc

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run(filename string) {
	inputBytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	inputString := string(inputBytes)
	value := DoCalc(inputString)
	fmt.Println(value)
}

func ParseInputString(input string) ([][]int, []string) {
	returnNumbers := [][]int{}
	rows := strings.Split(strings.TrimSpace(input), "\n")
	lastRowNumber := len(rows)-1
	// Take the operator from the last row
	returnOperators := strings.Fields(rows[lastRowNumber])

	// Skip the last row
	for _, row := range rows[:lastRowNumber] {
		items := strings.Fields(row)
		numbersInCurrentRow := []int{}
		for _, value := range items {
			number, _ := strconv.Atoi(value)
			numbersInCurrentRow = append(numbersInCurrentRow, number)
		}
		returnNumbers = append(returnNumbers, numbersInCurrentRow)
	}	
	return returnNumbers, returnOperators
}

func DoCalc(input string) int {
	returnValue := 0

	numbers, operators := ParseInputString(input)
	numberOfRows := len(numbers)

	// Each operator is a proxy for the column
	for operatorIndex, operatorValue := range operators {
		// start with the first number in row
		columnProduct := numbers[0][operatorIndex]
		// iterate for each number after
		for rowIndex := 1; rowIndex < numberOfRows; rowIndex++ {
			rowValue := numbers[rowIndex][operatorIndex]

			switch operatorValue {
			case "+":
				columnProduct += rowValue
			case "*":
				columnProduct *= rowValue
			}
		}
		returnValue += columnProduct
	}

	return returnValue
}