package aoc

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	// "slices"
	"strconv"
	"strings"
)

func IsInvalidId(inputVal int) bool {
	inputString := strconv.Itoa(inputVal)
	stringLength := len(inputString)

	if stringLength <= 1 {
		return false
	}

	// halfLength :=  stringLength / 2

	// asRunes := []rune(inputString)
	// var firstHalf = []rune{}
	// var secondHalf = []rune{}

	// if stringLength % 2 == 0 {
	// 	firstHalf = asRunes[:halfLength]
	// 	secondHalf = asRunes[halfLength:]	
	// } else {
	// 	firstHalf = asRunes[:halfLength+1]
	// 	secondHalf = asRunes[halfLength:]
	// }
	
	// return slices.Equal(firstHalf, secondHalf)

	for i := 1; i < len(inputString); i++ {
		subString := inputString[:i]
		currentPattern := ""
		for j := 0; len(currentPattern) < stringLength; j++ {
			currentPattern += subString
		}
		if currentPattern == inputString {
			return true
		}
	}

	return false
}

func ExpandRange(inputRange string) ([]int, error) {

	if !strings.Contains(inputRange, "-") {
		formattedErrMessage := fmt.Sprintf("invalid input: %s", inputRange)
		return nil, errors.New(formattedErrMessage)
	}

	index := strings.LastIndexByte(inputRange, '-')
	
	if index <= 0 {
		formattedErrMessage := fmt.Sprintf("invalid input: %s", inputRange)
		return nil, errors.New(formattedErrMessage)
	}
	rangeStart, _ := strconv.Atoi(inputRange[:index])

	if index +1 >= len(inputRange) {
		formattedErrMessage := fmt.Sprintf("invalid input: %s", inputRange)
		return nil, errors.New(formattedErrMessage)
	}
	rangeEnd, _ := strconv.Atoi(inputRange[index+1:])

	returnRange := []int{}
	for i := rangeStart; i <= rangeEnd; i++ {
		returnRange = append(returnRange, i)
	}

	return returnRange, nil
}

func ReturnListOfInvalidIdsFromRange(inputRange string) ([]int, error) {

	expandedRange, expandRangeError := ExpandRange(inputRange)
	if expandRangeError != nil {
		return nil, expandRangeError
	}

	invalidIds := []int{}
	// Gotcha: watch out for index vs value
	for _, id := range expandedRange {
		if IsInvalidId(id) {
			invalidIds = append(invalidIds, id)
		}
	}

	return invalidIds, nil
}

func SumRange(inputRange []int) int {
	sum := 0
	for _, value := range inputRange {
		sum += value
	}
	return sum
}

/*
	cd 2025/day-02/
	go run ./cmd/
*/

func Run() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	fileLines := []string{}
	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	allInOne := strings.Join(fileLines, "\n")
	splitRanges := strings.Split(allInOne, ",")

	accumulatedInvalidIds := []int{}
	for _, currentRange := range splitRanges {
		invalidIdsFound, err := ReturnListOfInvalidIdsFromRange(currentRange)
		if err != nil {
			panic(err)
		}
		accumulatedInvalidIds = append(accumulatedInvalidIds, invalidIdsFound...)
	}

	// fmt.Println(accumulatedInvalidIds)
	sum := SumRange(accumulatedInvalidIds)
	fmt.Println(sum)

}