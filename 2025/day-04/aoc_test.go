package aoc

import (
	"maps"
	"testing"
)

func TestCurrentCharacterIsPaperRollTrue(t *testing.T) {
	want := true
	// input := coord{0,0}
	// got := IsCurrentCoordPaperRoll(input)
	got := IsCurrentCharacterPaperRoll('@')
	if got != want {
		t.Errorf("Wanted %v, but got %v", want, got)		
	}
}

func TestCurrentCharacterIsPaperRollFalse(t *testing.T) {
	want := false
	got := IsCurrentCharacterPaperRoll('.')
	if got != want {
		t.Errorf("Wanted %v, but got %v", want, got)
	}
}

func TestStringToGrid(t *testing.T) {
	input :=   `...
				...
				...
				`
	want := make(map[coord]rune)
	want[coord{0,0}] = '.'
	want[coord{0,1}] = '.'
	want[coord{0,2}] = '.'
	want[coord{1,0}] = '.'
	want[coord{1,1}] = '.'
	want[coord{1,2}] = '.'
	want[coord{2,0}] = '.'
	want[coord{2,1}] = '.'
	want[coord{2,2}] = '.'
	got := ParseStringToGrid(input)
	if !maps.Equal(want, got) {
		t.Errorf("Wanted %v, but got %v", want, got)
	}
}

func TestGetNeighboursSafe(t *testing.T) {
	gridInput := `
	...
	...
	...
	`
	grid := ParseStringToGrid(gridInput)
	want := make(map[coord]rune)
	want[coord{0,0}] = '.'
	want[coord{0,1}] = '.'
	want[coord{0,2}] = '.'
	want[coord{1,0}] = '.'
	// want[coord{1,1}] = '.' // not our current position
	want[coord{1,2}] = '.'
	want[coord{2,0}] = '.'
	want[coord{2,1}] = '.'
	want[coord{2,2}] = '.'
	currentCoordinate := coord{1, 1}
	got := GetNeighbours(grid, currentCoordinate)
	if !maps.Equal(want, got) {
		t.Errorf("Wanted %v, but got %v", want, got)
	}
}

func TestGetNeighboursTop(t *testing.T) {
	gridInput := `
	...
	...
	...
	`
	grid := ParseStringToGrid(gridInput)
	want := make(map[coord]rune)
	want[coord{0,0}] = '.'
	//want[coord{0,1}] = '.'
	want[coord{0,2}] = '.'
	want[coord{1,0}] = '.'
	want[coord{1,1}] = '.'
	want[coord{1,2}] = '.'
	// want[coord{2,0}] = '.'
	// want[coord{2,1}] = '.'
	// want[coord{2,2}] = '.'
	currentCoordinate := coord{0, 1}
	got := GetNeighbours(grid, currentCoordinate)
	if !maps.Equal(want, got) {
		t.Errorf("Wanted %v, but got %v", want, got)
	}
}

func TestGetNeighboursRight(t *testing.T) {
	gridInput := `
	...
	...
	...
	`
	grid := ParseStringToGrid(gridInput)
	want := make(map[coord]rune)
	//want[coord{0,0}] = '.'
	want[coord{0,1}] = '.'
	want[coord{0,2}] = '.'
	//want[coord{1,0}] = '.'
	want[coord{1,1}] = '.'
	//want[coord{1,2}] = '.'
	// want[coord{2,0}] = '.'
	want[coord{2,1}] = '.'
	want[coord{2,2}] = '.'
	currentCoordinate := coord{1, 2}
	got := GetNeighbours(grid, currentCoordinate)
	if !maps.Equal(want, got) {
		t.Errorf("Wanted %v, but got %v", want, got)
	}
}

func TestGetNeighboursLeft(t *testing.T) {
	gridInput := `
	...
	...
	...
	`
	grid := ParseStringToGrid(gridInput)
	want := make(map[coord]rune)
	want[coord{0,0}] = '.'
	want[coord{0,1}] = '.'
	//want[coord{0,2}] = '.'
	//want[coord{1,0}] = '.'
	want[coord{1,1}] = '.'
	//want[coord{1,2}] = '.'
	want[coord{2,0}] = '.'
	want[coord{2,1}] = '.'
	//want[coord{2,2}] = '.'
	currentCoordinate := coord{1, 0}
	got := GetNeighbours(grid, currentCoordinate)
	if !maps.Equal(want, got) {
		t.Errorf("Wanted %v, but got %v", want, got)
	}
}

func TestGetNeighboursBottom(t *testing.T) {
	gridInput := `
	...
	...
	...
	`
	grid := ParseStringToGrid(gridInput)
	want := make(map[coord]rune)
	//want[coord{0,0}] = '.'
	//want[coord{0,1}] = '.'
	//want[coord{0,2}] = '.'
	want[coord{1,0}] = '.'
	want[coord{1,1}] = '.'
	want[coord{1,2}] = '.'
	want[coord{2,0}] = '.'
	//want[coord{2,1}] = '.'
	want[coord{2,2}] = '.'
	currentCoordinate := coord{2, 1}
	got := GetNeighbours(grid, currentCoordinate)
	if !maps.Equal(want, got) {
		t.Errorf("Wanted %v, but got %v", want, got)
	}
}

func TestUpdateGrid(t *testing.T){
	gridInput := `
	..@@.@@@@.
	@@@.@.@.@@
	@@@@@.@.@@
	@.@@@@..@.
	@@.@@@@.@@
	.@@@@@@@.@
	.@.@.@.@@@
	@.@@@.@@@@
	.@@@@@@@@.
	@.@.@@@.@.
	`
	wantInput := `
	..xx.xx@x.
	x@@.@.@.@@
	@@@@@.x.@@
	@.@@@@..@.
	x@.@@@@.@x
	.@@@@@@@.@
	.@.@.@.@@@
	x.@@@.@@@@
	.@@@@@@@@.
	x.x.@@@.x.
	`
	expectedGrid := ParseStringToGrid(wantInput)
	expectedNumberUpdated := 13
	grid := ParseStringToGrid(gridInput)
	updatedGrid, numberUpdated := UpdateGridWherePaperRollsCanBeMoved(grid)
	if !maps.Equal(expectedGrid, updatedGrid) || expectedNumberUpdated != numberUpdated {
		t.Errorf("Wanted \n %v \n with %v updates \n but got \n %v with %v updates", expectedGrid,expectedNumberUpdated, updatedGrid, numberUpdated)
	}
}
