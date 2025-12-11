package day06

import (
	"bufio"
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
	part2Value := DoCalcPart2(filename)
	fmt.Println(value, part2Value)
}

func ParseInputString(input string) ([][]int, []string) {
	returnNumbers := [][]int{}
	rows := strings.Split(strings.TrimSpace(input), "\n")
	lastRowNumber := len(rows) - 1
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

func Part2ParseInputFile(fileName string) ([][]int, []string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close() //nolint:errcheck
	// can't use fields as the blanks are important.

	scanner := bufio.NewScanner(file)
	returnNumbers := [][]int{}
	returnOperators := []string{}

	rows := []string{}
	for scanner.Scan() {
		currentLine := scanner.Text()
		rows = append(rows, currentLine)
	}

	// Want to start top row, but last colum
	// Use operator to step out of current equation?
	endColumnIndex := len(rows[0]) // 15
	currentEquation := []int{}
	for col := endColumnIndex - 1; col >= 0; col-- {
		builder := ""
		// isOperator := false

		for row := range rows {
			currentChar := rows[row][col]

			switch currentChar {
			case ' ':
				if builder != "" {
					conv, _ := strconv.Atoi(builder)
					// fmt.Printf("End of number %v\n", conv)
					currentEquation = append(currentEquation, conv)
					builder = "" // reset
				}
			case '+', '*':
				if builder != "" {
					// bank the current number
					conv, _ := strconv.Atoi(builder)
					currentEquation = append(currentEquation, conv)
					builder = "" // reset
				}

				// Add to return values
				returnOperators = append(returnOperators, string(currentChar))
				returnNumbers = append(returnNumbers, currentEquation)
				// fmt.Printf("End of equation %v with %v\n", currentEquation, string(currentChar))

				// reset for next block
				currentEquation = []int{} // reset
			default:
				// should be a digit
				builder += string(currentChar)
			}
			// fmt.Printf("Currently building %v\n", builder)
		}
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

func DoCalcPart2(fileName string) int {
	returnValue := 0
	numbers, operators := Part2ParseInputFile(fileName)

	for operatorIndex, operatorValue := range operators {
		numbersInRow := numbers[operatorIndex]

		// start with the first number in row
		product := numbersInRow[0]
		remainingNumbersInRow := numbersInRow[1:]
		// fmt.Println(remainingNumbersInRow)

		// iterate for each number after
		for _, nextNumber := range remainingNumbersInRow {

			// fmt.Printf("Calculating %v %v %v\n", product, operatorValue, nextNumber)

			switch operatorValue {
			case "+":
				product += nextNumber
			case "*":
				product *= nextNumber
			}
		}
		returnValue += product
	}
	return returnValue
}
