package day09

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func Run(fileName string) {
	inputBytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	tiles := GetListOfTilesFromString(string(inputBytes))
	rectangles := GetListOfRectanglesFromTiles(tiles)
	sort.Slice(rectangles, func(a, b int) bool { return rectangles[a].area > rectangles[b].area })
	largestArea := rectangles[0].area
	fmt.Println("Part 1: ", largestArea)

	greenRectangles := GetGreenRectangles(tiles)
	maxRect := GetLargestAreaOfAnInternalRectangle(rectangles, greenRectangles)
	fmt.Println("Part 2: ", maxRect)
}

type Tile struct {
	x, y int
}

type Rectangle struct {
	t1, t2 Tile
	area   int
}

func GetListOfTilesFromString(input string) []Tile {
	returnValue := []Tile{}

	for _, rowText := range strings.Split(string(input), "\n") {
		if len(strings.Fields(rowText)) == 0 {
			continue // step over blank rows
		}
		var x, y int
		_, err := fmt.Sscanf(rowText, "%d,%d", &x, &y)
		if err != nil {
			panic(err)
		}
		returnValue = append(returnValue, Tile{x, y})
	}

	return returnValue
}

func CalculateArea(a, b Tile) int {
	return (max(a.x, b.x) - min(a.x, b.x) + 1) * (max(a.y, b.y) - min(a.y, b.y) + 1)
}

func GetListOfRectanglesFromTiles(tiles []Tile) []Rectangle {
	returnValue := []Rectangle{}

	for i := 0; i < len(tiles)-1; i++ {
		for j := i + 1; j < len(tiles); j++ {
			tileA := tiles[i]
			tileB := tiles[j]
			areaValue := CalculateArea(tileA, tileB)
			returnValue = append(returnValue, Rectangle{tileA, tileB, areaValue})
		}
	}

	return returnValue
}

func GetGreenRectangles(tiles []Tile) []Rectangle {
	returnValue := []Rectangle{}

	n := len(tiles)
	for i := 0; i < n; i++ {

		tileA := tiles[i]
		tileB := tiles[(i+1)%n]
		// fmt.Printf("TileA %v and TileB %v\n", tileA, tileB)

		area := CalculateArea(tileA, tileB)

		// I don't think we're interested in the area for the green bits
		returnValue = append(returnValue, Rectangle{tileA, tileB, area})
	}

	return returnValue
}

func CheckIntersect(segment Rectangle, t1, t2 Tile) bool {
	minX := min(t1.x, t2.x) + 1
	maxX := max(t1.x, t2.x) - 1
	minY := min(t1.y, t2.y) + 1
	maxY := max(t1.y, t2.y) - 1

	segMinX := min(segment.t1.x, segment.t2.x)
	segMaxX := max(segment.t1.x, segment.t2.x)
	segMinY := min(segment.t1.y, segment.t2.y)
	segMaxY := max(segment.t1.y, segment.t2.y)

	if (segMaxX < minX || segMinX > maxX) || (segMaxY < minY || segMinY > maxY) {
		return false
	}
	return true
}

func GetLargestAreaOfAnInternalRectangle(redRectangles, greenRectangles []Rectangle) int {
	maxRect := 0

outer:
	for _, rect := range redRectangles {
		if rect.area < maxRect {
			continue // skip over small rectangles
		}
		for _, greenRect := range greenRectangles {
			if CheckIntersect(greenRect, rect.t1, rect.t2) {
				continue outer
			}
		}
		// fmt.Printf("We got one %v\n", rect)
		maxRect = rect.area
	}

	return maxRect
}
