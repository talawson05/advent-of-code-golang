package aoc

import (
	"errors"
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

func main() {

}