package aoc

import (
	"fmt"
	"os"
	"strings"
)

func Run(filename string) {
	inputBytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	inputString := string(inputBytes)
	grid := ParseStringToGrid(inputString)
	_, numberUpdated := UpdateGridWherePaperRollsCanBeMoved(grid)
	fmt.Println(numberUpdated)	
}

type coord struct {
	x int
	y int
}

func IsCurrentCharacterPaperRoll(input rune) bool {
	return input == '@' || input == 'x'
}

func ParseStringToGrid(input string) map[coord]rune {
	returnValue := make(map[coord]rune)
	for row, rowString := range strings.Fields(input) {
		for column, character := range rowString {
			returnValue[coord{row, column}] = character
		}
	}

	return returnValue
}

func GetNeighbours(grid map[coord]rune, currentCoordinate coord) map[coord]rune {
	x := currentCoordinate.x
	y := currentCoordinate.y
	goodNeighbours := make(map[coord]rune) // Everybody needs good neeeeeeeiiiighbours

	for row := x -1; row < x + 2; row++ {
		for column := y-1; column < y+2; column++ {
			coordToCheck := coord{row, column}
			if coordToCheck == currentCoordinate {
				// can't be a neighbour to yourself
				continue
			}

			if inBounds(grid, coordToCheck) {
				// store the neighbour coords and value
				goodNeighbours[coordToCheck] = grid[coordToCheck]
			}
		}
	}
	return goodNeighbours
}

func inBounds(grid map[coord]rune, currentCord coord) bool {
	_, ok := grid[currentCord]
	return ok
}

func UpdateGridWherePaperRollsCanBeMoved(grid map[coord]rune) (map[coord]rune, int) {
	numberOfUpdatedCells := 0

	for cell := range grid {

		if IsCurrentCharacterPaperRoll(grid[cell]) {
			// Current cell is @

			neighbours := GetNeighbours(grid, cell)
			// How many neighbours have value @
			neighbouringPaperRolls := 0
			for n := range neighbours {
				if IsCurrentCharacterPaperRoll(neighbours[n]) {
					neighbouringPaperRolls++
				}
			}

			if neighbouringPaperRolls < 4 {
				grid[cell] = 'x'
				numberOfUpdatedCells++
			}
		}
	}

	// Return the updated grid and how many updates were made
	return grid, numberOfUpdatedCells
}