package aoc

import (
	"fmt"
	"sort"
	"testing"
)

func TestGetListOfTilesFromString(t *testing.T) {
	input := `
	7,1
	11,1
	11,7
	9,7
	9,5
	2,5
	2,3
	7,3
	`
	wantLength := 8
	expectedTile := Tile{9, 7}
	got := GetListOfTilesFromString(input)

	if wantLength != len(got) {
		t.Errorf("Wanted %v but got %v", wantLength, got)
	}

	if got[3] != expectedTile {
		t.Errorf("Wanted %v in %v", expectedTile, got)
	}

	fmt.Println("foo")
}

func TestCalculateArea(t *testing.T) {
	a := Tile{2, 5}
	b := Tile{9, 7}
	want := 24
	got := CalculateArea(a, b)
	if want != got {
		t.Errorf("want %v but got %v", want, got)
	}
}

func TestCalculateAreaLine(t *testing.T) {
	a := Tile{7, 3}
	b := Tile{2, 3}
	want := 6
	got := CalculateArea(a, b)
	if want != got {
		t.Errorf("want %v but got %v", want, got)
	}
}

func TestGetListOfRectanglesFromTiles(t *testing.T) {
	input := `
	7,1
	11,1
	11,7
	9,7
	9,5
	2,5
	2,3
	7,3
	`
	tiles := GetListOfTilesFromString(input)
	got := GetListOfRectanglesFromTiles(tiles)
	wantLength := 28
	expectedRectangle := Rectangle{
		Tile{9, 5},
		Tile{2, 5},
		8,
	}

	if wantLength != len(got) {
		t.Errorf("Wanted %v but got %v", wantLength, got)
	}

	if got[22] != expectedRectangle {
		t.Errorf("Wanted %v in %v", expectedRectangle, got)
	}

	sort.Slice(got, func(a, b int) bool { return got[a].area > got[b].area })
	largestArea := got[0].area
	want := 50
	if want != largestArea {
		t.Errorf("want %v but got %v", want, largestArea)
	}
}

func TestGetGreenRectangles(t *testing.T) {
	input := `
	7,1
	11,1
	11,7
	9,7
	9,5
	2,5
	2,3
	7,3
	`
	tiles := GetListOfTilesFromString(input)
	got := GetGreenRectangles(tiles)
	fmt.Println(len(got))
	fmt.Println(got)
}

func TestCheckIntersectTrue(t *testing.T) {
	rectangle := Rectangle{
		Tile{11, 1},
		Tile{2, 5},
		50,
	}
	tileA := Tile{11, 1}
	tileB := Tile{7, 3}

	got := CheckIntersect(rectangle, tileA, tileB)
	want := true

	if want != got {
		t.Errorf("Wanted %v but got %v", want, got)
	}
}

func TestCheckIntersectFalse(t *testing.T) {
	rectangle := Rectangle{
		Tile{11, 1},
		Tile{2, 5},
		50,
	}
	tileA := Tile{0, 0}
	tileB := Tile{6, 1}

	got := CheckIntersect(rectangle, tileA, tileB)
	want := false

	if want != got {
		t.Errorf("Wanted %v but got %v", want, got)
	}
}

func TestGetLargestAreaOfAnInternalRectangle(t *testing.T) {
	input := `
	7,1
	11,1
	11,7
	9,7
	9,5
	2,5
	2,3
	7,3
	`
	tiles := GetListOfTilesFromString(input)
	redRectangles := GetListOfRectanglesFromTiles(tiles)
	greenRectangles := GetGreenRectangles(tiles)
	got := GetLargestAreaOfAnInternalRectangle(redRectangles, greenRectangles)
	want := 24
	if want != got {
		t.Errorf("want %v but got %v", want, got)
	}
}
