package aoc

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func IsInvalidId(inputVal int) bool {
	inputString := strconv.Itoa(inputVal)
	stringLength := len(inputString)
	halfLength :=  stringLength / 2

	if (stringLength % 2) != 0 {
		return false
	}

	asRunes := []rune(inputString)
	firstHalf := asRunes[:halfLength]
	secondHalf := asRunes[halfLength:]	
	
	return slices.Equal(firstHalf, secondHalf)
}

func ExpandRange(inputRange string) ([]int, error) {

	if !strings.Contains(inputRange, "-") {
		formattedErrMessage := fmt.Sprintf("invalid input: %s", inputRange)
		return nil, errors.New(formattedErrMessage)
	}

	split := strings.Split(inputRange, "-")

	if len(split) != 2 {
		formattedErrMessage := fmt.Sprintf("invalid input: %s", inputRange)
		return nil, errors.New(formattedErrMessage)
	}

	returnRange := []int{}
	rangeStart, _ := strconv.Atoi(split[0])
	rangeEnd, _ := strconv.Atoi(split[1])
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

	fmt.Println(accumulatedInvalidIds)
	sum := SumRange(accumulatedInvalidIds)
	fmt.Println(sum)

}