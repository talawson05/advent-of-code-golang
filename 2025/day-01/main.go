package aoc

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func DoRotation(startingPosition int, direction string, rotationValue int) (int, error) {
	currentPosition := startingPosition

	for counter:= 0; counter <rotationValue ; counter ++ {

		switch {
		case direction == "L":
			currentPosition--
		case direction == "R":
			currentPosition++
		default:
			return -1, errors.New("unexpected direction")
		}
		
		if currentPosition > 99 {
			currentPosition = 0			
		} else if currentPosition < 0 {
			currentPosition = 99			
		}
	}

	return currentPosition, nil
}

func Run() {
	currentDialPosition := 50
	countOfLandingOnZero := 0
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()
		stepDirection := currentLine[:1]
		stepClicks, stepParseErr := strconv.Atoi(currentLine[1:])
		if stepParseErr != nil {
			panic(stepParseErr)
		}
		returnPosition, stepErr := DoRotation(currentDialPosition, stepDirection, stepClicks)
		currentDialPosition = returnPosition
		if stepErr != nil {
			panic(stepErr)
		}
		if currentDialPosition == 0 {
			countOfLandingOnZero++		
		}
	}
	fmt.Printf("Final dial position: %d\n", currentDialPosition)
	fmt.Printf("Number of times we landed on 0: %d\n", countOfLandingOnZero)

}