package aoc

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Run(filename string) {
	inputBytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	inputString := string(inputBytes)
	count, sum := CountOfFreshIngredients(inputString)
	fmt.Println(count)
	fmt.Println(sum)
}

func ExpandRange(inputRange string) ([]int, error) {

	rangeStart, rangeEnd, err := GetLowAndHighFromRange(inputRange)
	if err != nil {
		return nil, err
	}

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
	rangeStart, rangeEnd, err := GetLowAndHighFromRange(inputRange)
	if err != nil {
		return false, err
	}

	return target >= rangeStart && target <= rangeEnd, nil
}

func ItemIsInSet(inputSet map[int]bool, targetItem int) bool {
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

func GetLowAndHighFromRange(inputRange string) (int, int, error) {
	if !strings.Contains(inputRange, "-") {
		formattedErrMessage := fmt.Sprintf("invalid input: %s", inputRange)
		return -1, -1, errors.New(formattedErrMessage)
	}

	index := strings.LastIndexByte(inputRange, '-')

	if index <= 0 {
		formattedErrMessage := fmt.Sprintf("invalid input: %s", inputRange)
		return -1, -1, errors.New(formattedErrMessage)
	}
	rangeStart, _ := strconv.Atoi(inputRange[:index])

	if index+1 >= len(inputRange) {
		formattedErrMessage := fmt.Sprintf("invalid input: %s", inputRange)
		return -1, -1, errors.New(formattedErrMessage)
	}
	rangeEnd, _ := strconv.Atoi(inputRange[index+1:])

	return rangeStart, rangeEnd, nil
}

func CountOfFreshIngredients(input string) (int, int) {
	returnCount := 0
	ranges, IDs := ParseInput(input)

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

	// Part 2

	mergedRanges := MergeRanges(ranges)
	sum := CountTotalInRanges(mergedRanges)

	return returnCount, sum
}

func MergeRanges(inputRanges []string) []string {
	mergedRanges := []string{}

	// we need the ranges in order so we can iterate over and look for overlaps
	sort.Slice(inputRanges, func(i, j int) bool {
		iLow, _, _ := GetLowAndHighFromRange(inputRanges[i])
		jennyFromTheBlock, _, _ := GetLowAndHighFromRange(inputRanges[j])
		return iLow < jennyFromTheBlock
	})

	previousRange := inputRanges[0]
	for i := 1; i < len(inputRanges); i++ {

		previousLow, previousHigh, _ := GetLowAndHighFromRange(previousRange)
		currentRange := inputRanges[i]
		currentLow, currentHigh, _ := GetLowAndHighFromRange(currentRange)

		// 10-14 & 12-18
		// 12 <= 14
		if currentLow <= previousHigh {
			// We have an overlap!

			// 18 > 14
			if currentHigh > previousHigh {
				// We've got a new upper bound, overwrite it
				previousRange = fmt.Sprintf("%d-%d", previousLow, currentHigh)
			}
		} else {
			mergedRanges = append(mergedRanges, previousRange)
			previousRange = currentRange
		}
	}
	mergedRanges = append(mergedRanges, previousRange)

	return mergedRanges
}

func CountTotalInRanges(inputRanges []string) int {
	returnCount := 0

	for _, currentRange := range inputRanges {
		low, high, _ := GetLowAndHighFromRange(currentRange)
		// 5 - (3-1)
		count := high - (low - 1)
		returnCount += count
	}

	return returnCount
}
