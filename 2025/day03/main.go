package aoc

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	defer file.Close()

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

func GetBiggestNumberFromRange(input string) string {
	targetLength := 12
	sliceOfInts := stringToSliceOfInts(input)
	returnSliceOfInts := make([]int, targetLength) // all 0s

	for outerIndex, outerInt := range sliceOfInts {
		for innerIndex, innerInt := range returnSliceOfInts {
			// a - b > 15 - 12
			if outerIndex-innerIndex > len(sliceOfInts)-targetLength {
				// We're too far along the string, and won't make 12 digits
				continue // increment inner loop
			}

			if outerInt > innerInt {
				// Set new bigger number
				returnSliceOfInts[innerIndex] = outerInt
				break // increment outer loop
			}
		}
	}

	return concatIntsToString(returnSliceOfInts)
}

func stringToSliceOfInts(input string) []int {
	sliceOfInts := make([]int, 0, len(input))
	for _, character := range input {
		// note: convert rune to int - https://stackoverflow.com/questions/21322173/convert-rune-to-int#comment112055881_21322694
		sliceOfInts = append(sliceOfInts, int(character-'0'))
	}

	return sliceOfInts
}

func concatIntsToString(ints []int) string {
	var builder strings.Builder
	for _, v := range ints {
		builder.WriteString(strconv.Itoa(v))
	}

	return builder.String()
}
