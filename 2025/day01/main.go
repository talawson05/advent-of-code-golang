package day01

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func DoRotation(startingPosition int, direction string, rotationValue int, zeroCounter ...int) (int, int, error) {
	currentPosition := startingPosition

	for counter := 0; counter < rotationValue; counter++ {

		switch direction {
		case "L":
			currentPosition--
		case "R":
			currentPosition++
		default:
			return -1, -1, errors.New("unexpected direction")
		}

		if currentPosition > 99 {
			currentPosition = 0
		} else if currentPosition < 0 {
			currentPosition = 99
		}

		if len(zeroCounter) > 0 && currentPosition == 0 {
			zeroCounter[0]++
		}
	}

	if len(zeroCounter) > 0 {
		return currentPosition, zeroCounter[0], nil
	}
	return currentPosition, 0, nil
}

/*
	cd 2025/day-01/
	go run ./cmd/
*/

func Run() {
	currentDialPosition := 50
	countOfLandingOnZero := 0
	countOfTouchingZero := 0
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
		// GOTCHA: occlusion
		returnPosition, returnZeroCount, stepErr := DoRotation(currentDialPosition, stepDirection, stepClicks, countOfTouchingZero)
		currentDialPosition = returnPosition
		countOfTouchingZero = returnZeroCount
		if stepErr != nil {
			panic(stepErr)
		}
		if currentDialPosition == 0 {
			countOfLandingOnZero++
		}
	}
	fmt.Printf("Final dial position: %d\n", currentDialPosition)
	fmt.Printf("Number of times we landed on 0: %d\n", countOfLandingOnZero)
	fmt.Printf("Number of times we tocuhed 0: %d\n", countOfTouchingZero)
}
