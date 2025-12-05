package aoc

import (
	"errors"
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
	count := CountOfFreshIngredients(inputString)
	fmt.Println(count)
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

	fmt.Printf("Expanding range between %d and %d\n", rangeStart, rangeEnd)
	returnRange := []int{}
	for i := rangeStart; i <= rangeEnd; i++ {
		returnRange = append(returnRange, i)
	}

	return returnRange, nil
}

func CombineRangesIntoSet(inputListOfRanges [][]int) map[int]bool {
	returnMap := map[int]bool{}
	for _, currentRange := range inputListOfRanges {
		for _, value := range currentRange {
			returnMap[value] = true
		}
	}

	return returnMap
}

func ItemIsInRange(inputRange string, target int) (bool, error) {
	if !strings.Contains(inputRange, "-") {
		formattedErrMessage := fmt.Sprintf("invalid input: %s", inputRange)
		return false, errors.New(formattedErrMessage)
	}

	index := strings.LastIndexByte(inputRange, '-')
	
	if index <= 0 {
		formattedErrMessage := fmt.Sprintf("invalid input: %s", inputRange)
		return false, errors.New(formattedErrMessage)
	}
	rangeStart, _ := strconv.Atoi(inputRange[:index])

	if index +1 >= len(inputRange) {
		formattedErrMessage := fmt.Sprintf("invalid input: %s", inputRange)
		return false, errors.New(formattedErrMessage)
	}
	rangeEnd, _ := strconv.Atoi(inputRange[index+1:])

	return target >= rangeStart && target <= rangeEnd, nil
}

func ItemIsInSet(inputSet  map[int]bool, targetItem int) bool {
	_, ok := inputSet[targetItem]
	return ok
}

func ParseInput(input string) ([]string, []int) {
	returnRanges := []string{}
	returnIds := []int{}

	for _, rowString := range strings.Fields(input) {

		if strings.Contains(rowString, "-") {
			returnRanges = append(returnRanges, rowString)
		} else {
			converted, convertError := strconv.Atoi(rowString)
			if convertError != nil {
				formattedErrMessage := fmt.Sprintf("unable to convert ID to int: %v", rowString)
				panic(errors.New(formattedErrMessage))
			}
			returnIds = append(returnIds, converted)
		}
		// Blank row is filtered out by the strings.fields
	}

	return returnRanges, returnIds
}

func CountOfFreshIngredients(input string) int {
	returnCount := 0

	ranges, IDs := ParseInput(input)
	fmt.Println("input parsed")
	fmt.Printf("Number of ranges is %v\n", len(ranges))
	fmt.Printf("Number of IDs is %v\n", len(IDs))

	for _, id := range IDs {
		for _, currentRange := range ranges {
			itemInRange, errCheckingRange := ItemIsInRange(currentRange, id)
			if errCheckingRange != nil {
				formattedErrMessage := fmt.Sprintf("error when checking item in range: %v, %v", id, currentRange)
				panic(errors.New(formattedErrMessage))
			}
			if itemInRange {
				// fmt.Printf("Id %v found in range %v\n", id, currentRange)
				returnCount++
				break // increment the outer loop to avoid the duplicates
			}
		}
	}

	return returnCount
}
