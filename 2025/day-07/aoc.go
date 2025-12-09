package aoc

import (
	"fmt"
	"os"
	"strings"
)

func Run(fileName string) {
	allText := ReadAllTextFromFile(fileName)
	numberOfSplits, numberOfTimelines := TrackSplitsOnGrid(allText)
	fmt.Println(numberOfSplits)
	fmt.Println(numberOfTimelines)
}

func ReadAllTextFromFile(fileName string) string {
	inputBytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(inputBytes)
}

func TrackSplitsOnGrid(input string) (int, int) {

	allText := strings.Fields(input)

	beams := make([]int, len(allText[0]))

	// Start the count of beams for the first row of input
	beams[strings.Index(allText[0], "S")] = 1

	splitCounter, timelineCounter := 0, 1
	for _, rowText := range allText {
		for columnIndex, characterRune := range rowText {
			if characterRune == '^' {

				valueInCurrentPosition := beams[columnIndex]
				// because not every splitter receives a beam, step over
				if valueInCurrentPosition == 0 {
					continue
				}
				// increment the timelines, note we want both where L/R were merged
				timelineCounter += valueInCurrentPosition

				splitCounter += 1
				// Increment the positions either side of the split
				beams[columnIndex-1] += valueInCurrentPosition
				beams[columnIndex+1] += valueInCurrentPosition

				// Set the current position to blank as beam is now on either side
				beams[columnIndex] = 0
			}
		}
	}

	return splitCounter, timelineCounter
}