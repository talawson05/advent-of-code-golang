package day04

import (
	"maps"
	"testing"
)

func TestCurrentCharacterIsPaperRollTrue(t *testing.T) {
	want := true
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
	input := `...
				...
				...
				`
	want := map[coord]rune{
		{0, 0}: '.',
		{0, 1}: '.',
		{0, 2}: '.',
		{1, 0}: '.',
		{1, 1}: '.',
		{1, 2}: '.',
		{2, 0}: '.',
		{2, 1}: '.',
		{2, 2}: '.',
	}
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
	want := map[coord]rune{
		{0, 0}: '.',
		{0, 1}: '.',
		{0, 2}: '.',
		{1, 0}: '.',
		// {1,1} : '.', // not our current position
		{1, 2}: '.',
		{2, 0}: '.',
		{2, 1}: '.',
		{2, 2}: '.',
	}
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
	want := map[coord]rune{
		{0, 0}: '.',
		// {0,1} : '.',
		{0, 2}: '.',
		{1, 0}: '.',
		{1, 1}: '.',
		{1, 2}: '.',
		// {2,0} : '.',
		// {2,1} : '.',
		// {2,2} : '.',
	}
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
	want := map[coord]rune{
		// {0,0} : '.',
		{0, 1}: '.',
		{0, 2}: '.',
		// {1,0} : '.',
		{1, 1}: '.',
		// {1,2} : '.',
		// {2,0} : '.',
		{2, 1}: '.',
		{2, 2}: '.',
	}
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
	want := map[coord]rune{
		{0, 0}: '.',
		{0, 1}: '.',
		// {0,2} : '.',
		// {1,0} : '.',
		{1, 1}: '.',
		// {1,2} : '.',
		{2, 0}: '.',
		{2, 1}: '.',
		// {2,2} : '.',
	}
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
	want := map[coord]rune{
		// {0,0} : '.',
		// {0,1} : '.',
		// {0,2} : '.',
		{1, 0}: '.',
		{1, 1}: '.',
		{1, 2}: '.',
		{2, 0}: '.',
		// {2,1} : '.',
		{2, 2}: '.',
	}
	currentCoordinate := coord{2, 1}
	got := GetNeighbours(grid, currentCoordinate)
	if !maps.Equal(want, got) {
		t.Errorf("Wanted %v, but got %v", want, got)
	}
}

func TestUpdateGrid(t *testing.T) {
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
		t.Errorf("Wanted \n %v \n with %v updates \n but got \n %v with %v updates", expectedGrid, expectedNumberUpdated, updatedGrid, numberUpdated)
	}
}

func TestRemoveUpdatesFromGrid(t *testing.T) {
	stringInput := `
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

	wantInput := `
	.......@..
	.@@.@.@.@@
	@@@@@...@@
	@.@@@@..@.
	.@.@@@@.@.
	.@@@@@@@.@
	.@.@.@.@@@
	..@@@.@@@@
	.@@@@@@@@.
	....@@@...
	`

	gridInput := ParseStringToGrid(stringInput)
	want := ParseStringToGrid(wantInput)
	got := RemoveUpdatesFromGrid(gridInput)
	if !maps.Equal(want, got) {
		t.Errorf("Wanted %v, but got %v", want, got)
	}

}

func TestUpdateGridRecursive(t *testing.T) {
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
	..........
	..........
	..........
	....@@....
	...@@@@...
	...@@@@@..
	...@.@.@@.
	...@@.@@@.
	...@@@@@..
	....@@@...
	`
	expectedGrid := ParseStringToGrid(wantInput)
	expectedNumberUpdated := 43
	grid := ParseStringToGrid(gridInput)
	updatedGrid, numberUpdated := RecursiveUpdateGridWherePaperRollsCanBeMoved(grid)
	if !maps.Equal(expectedGrid, updatedGrid) || expectedNumberUpdated != numberUpdated {
		t.Errorf("Wanted \n %v \n with %v updates \n but got \n %v with %v updates", expectedGrid, expectedNumberUpdated, updatedGrid, numberUpdated)
	}
}
