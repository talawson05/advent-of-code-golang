package day12

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run(fileName string) {
	presents, treeRegions := ParseInput(fileName)
	part1 := SolvePart1(presents, treeRegions)
	fmt.Println("Part1: ", part1)
}

type Coord struct {
	xPos, yPos int
}

type Present struct {
	id    int
	shape map[Coord]int
}

func (p Present) area() int {
	return len(p.shape)
}

type TreeRegion struct {
	width, height int
	scheme        []int
}

func (tr *TreeRegion) area() int {
	return tr.width * tr.height
}

func ParseInput(fileName string) ([]Present, []TreeRegion) {
	returnPresents := []Present{}
	returnTreeRegions := []TreeRegion{}

	raw, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	text := string(raw)
	lines := strings.Split(text, "\n")

	lineIndex := 0
	currentShapeRowIndex := 0
	present := Present{-1, map[Coord]int{}}
	for !strings.Contains(lines[lineIndex], "x") {
		if strings.Contains(lines[lineIndex], ":") {
			// Id line
			idString := strings.Split(lines[lineIndex], ":")[0]
			present.id, err = strconv.Atoi(idString)
			if err != nil {
				panic(err)
			}
		} else if len(lines[lineIndex]) == 0 {
			// blank, end of current preset section
			returnPresents = append(returnPresents, present)
			currentShapeRowIndex = 0 // reset for next shape
			present.id = -1          // defensive
			present.shape = map[Coord]int{}
		} else {
			// Present shape
			for characterPos, character := range lines[lineIndex] {
				switch character {
				case '#':
					present.shape[Coord{currentShapeRowIndex, characterPos}] = present.id
				case '.':
					// do nothing
				}
			}
			currentShapeRowIndex++

		}
		lineIndex++
	}

	for regionLineIndex := lineIndex; regionLineIndex < len(lines); regionLineIndex++ {
		x := strings.Index(lines[regionLineIndex], "x")
		width, err := strconv.Atoi(lines[regionLineIndex][:x])
		if err != nil {
			panic(err)
		}
		col := strings.Index(lines[regionLineIndex], ":")
		height, err := strconv.Atoi(lines[regionLineIndex][x+1 : col])
		if err != nil {
			panic(err)
		}

		parts := strings.Split(lines[regionLineIndex][col+2:], " ")
		presents := make([]int, 0, len(parts))
		for _, part := range parts {
			count, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}
			presents = append(presents, count)
		}
		currentTreeRegion := TreeRegion{width: width, height: height, scheme: presents}
		returnTreeRegions = append(returnTreeRegions, currentTreeRegion)
	}

	return returnPresents, returnTreeRegions
}

func SolvePart1(presents []Present, regions []TreeRegion) int {
	sum := 0
	for _, region := range regions {
		if TreeRegionAreaGreaterThanRequiredPresentsArea(region, presents) {
			sum++
		}
	}
	return sum
}

func TreeRegionAreaGreaterThanRequiredPresentsArea(treeRegion TreeRegion, presents []Present) bool {
	regionArea := treeRegion.area()
	presentsArea := 0
	for presentId, presentCount := range treeRegion.scheme {
		presentsArea += presentCount * presents[presentId].area()
	}
	return regionArea > presentsArea
}

func (currentShape *Present) RotatePresent() Present {
	returnValue := Present{currentShape.id, map[Coord]int{}}

	if value, ok := currentShape.shape[Coord{0, 0}]; ok {
		returnValue.shape[Coord{0, 2}] = value
	}
	if value, ok := currentShape.shape[Coord{0, 1}]; ok {
		returnValue.shape[Coord{1, 2}] = value
	}
	if value, ok := currentShape.shape[Coord{0, 2}]; ok {
		returnValue.shape[Coord{2, 2}] = value
	}
	if value, ok := currentShape.shape[Coord{1, 0}]; ok {
		returnValue.shape[Coord{0, 1}] = value
	}
	if value, ok := currentShape.shape[Coord{1, 1}]; ok {
		returnValue.shape[Coord{1, 1}] = value // mid point
	}
	if value, ok := currentShape.shape[Coord{1, 2}]; ok {
		returnValue.shape[Coord{2, 1}] = value
	}
	if value, ok := currentShape.shape[Coord{2, 0}]; ok {
		returnValue.shape[Coord{0, 0}] = value
	}
	if value, ok := currentShape.shape[Coord{2, 1}]; ok {
		returnValue.shape[Coord{1, 0}] = value
	}
	if value, ok := currentShape.shape[Coord{2, 2}]; ok {
		returnValue.shape[Coord{2, 0}] = value
	}

	return returnValue
}

func (currentShape *Present) FlipPresent() Present {
	returnValue := Present{currentShape.id, map[Coord]int{}}

	/*
		###
		.#.
		###
	*/
	if value, ok := currentShape.shape[Coord{0, 0}]; ok {
		returnValue.shape[Coord{0, 2}] = value
	}
	if value, ok := currentShape.shape[Coord{0, 1}]; ok {
		returnValue.shape[Coord{0, 1}] = value // unchanged
	}
	if value, ok := currentShape.shape[Coord{0, 2}]; ok {
		returnValue.shape[Coord{0, 0}] = value
	}
	if value, ok := currentShape.shape[Coord{1, 0}]; ok {
		returnValue.shape[Coord{1, 2}] = value
	}
	if value, ok := currentShape.shape[Coord{1, 1}]; ok {
		returnValue.shape[Coord{1, 1}] = value // mid point
	}
	if value, ok := currentShape.shape[Coord{1, 2}]; ok {
		returnValue.shape[Coord{1, 0}] = value
	}
	if value, ok := currentShape.shape[Coord{2, 0}]; ok {
		returnValue.shape[Coord{2, 2}] = value
	}
	if value, ok := currentShape.shape[Coord{2, 1}]; ok {
		returnValue.shape[Coord{2, 1}] = value // unchanged
	}
	if value, ok := currentShape.shape[Coord{2, 2}]; ok {
		returnValue.shape[Coord{2, 0}] = value
	}

	return returnValue
}
