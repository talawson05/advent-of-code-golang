package day12

import (
	"maps"
	"slices"
	"testing"
)

func TestParseInput(t *testing.T) {
	fileName := "example_input.txt"
	gotPresents, gotTreeRegions := ParseInput(fileName)
	wantPresentLength := 6
	/*
		###
		#..
		###
	*/
	wantSamplePresent := Present{
		4,
		map[Coord]int{
			Coord{0, 0}: 4,
			Coord{0, 1}: 4,
			Coord{0, 2}: 4,
			Coord{1, 0}: 4,
			Coord{2, 0}: 4,
			Coord{2, 1}: 4,
			Coord{2, 2}: 4,
		},
	}
	wantTreeRegionLength := 3
	// 12x5: 1 0 1 0 3 2
	wantSampleRegion := TreeRegion{
		12, 5, []int{1, 0, 1, 0, 3, 2},
	}

	if wantPresentLength != len(gotPresents) {
		t.Errorf("Wanted present length %v but got %v", wantPresentLength, len(gotPresents))
	}

	if wantTreeRegionLength != len(gotTreeRegions) {
		t.Errorf("Wanted region length %v but got %v", wantTreeRegionLength, len(gotTreeRegions))
	}

	if !maps.Equal(gotPresents[4].shape, wantSamplePresent.shape) {
		t.Errorf("Want %v but got %v", wantSamplePresent.shape, gotPresents[4].shape)
	}

	if gotTreeRegions[2].width != wantSampleRegion.width {
		t.Errorf("Want %v but got %v", wantSampleRegion.width, gotTreeRegions[2].width)
	}

	if gotTreeRegions[2].height != wantSampleRegion.height {
		t.Errorf("Want %v but got %v", wantSampleRegion.height, gotTreeRegions[2].height)
	}

	if !slices.Equal(gotTreeRegions[2].scheme, wantSampleRegion.scheme) {
		t.Errorf("Want %v but got %v", wantSampleRegion.scheme, gotTreeRegions[2].scheme)
	}
}

func TestPresentAreaCalc(t *testing.T) {
	inputPresent := Present{
		4,
		map[Coord]int{
			Coord{0, 0}: 4,
			Coord{0, 1}: 4,
			Coord{0, 2}: 4,
			Coord{1, 0}: 4,
			Coord{2, 0}: 4,
			Coord{2, 1}: 4,
			Coord{2, 2}: 4,
		},
	}
	want := 7
	got := inputPresent.area()
	if want != got {
		t.Errorf("wanted %v but got %v", want, got)
	}
}

func TestTreeRegionAreaCalc(t *testing.T) {
	input := TreeRegion{
		12, 5, []int{1, 0, 1, 0, 3, 2},
	}
	want := 60
	got := input.area()
	if want != got {
		t.Errorf("wanted %v but got %v", want, got)
	}
}

func TestTreeRegionAreaGreaterThanRequiredPresentsAreaTrue(t *testing.T) {
	fileName := "example_input.txt"
	presents, _ := ParseInput(fileName)
	exampleRegion := TreeRegion{
		12, 5, []int{1, 0, 1, 0, 3, 2},
	}
	got := TreeRegionAreaGreaterThanRequiredPresentsArea(exampleRegion, presents)
	if got != true {
		t.Errorf("Expected %v to be true", got)
	}
}

func TestTreeRegionAreaGreaterThanRequiredPresentsAreaFalse(t *testing.T) {
	fileName := "example_input.txt"
	presents, _ := ParseInput(fileName)
	exampleRegion := TreeRegion{
		2, 2, []int{1, 0, 1, 0, 3, 2},
	}
	got := TreeRegionAreaGreaterThanRequiredPresentsArea(exampleRegion, presents)
	if got != false {
		t.Errorf("Expected %v to be false", got)
	}
}

func TestSolvePart1(t *testing.T) {
	t.Skip("No longer required")
	fileName := "example_input.txt"
	presents, treeRegions := ParseInput(fileName)
	got := SolvePart1(presents, treeRegions)
	want := 2
	if want != got {
		t.Errorf("wanted %v but got %v", want, got)
	}
}

func TestRotatePresent(t *testing.T) {
	inputPresent := Present{
		4,
		map[Coord]int{
			Coord{0, 0}: 4,
			Coord{0, 1}: 4,
			Coord{0, 2}: 4,
			Coord{1, 0}: 4,
			Coord{2, 0}: 4,
			Coord{2, 1}: 4,
			Coord{2, 2}: 4,
		},
	}

	/*
		###
		#..
		###

		###
		#.#
		#.#
	*/
	wantPresent1 := Present{
		4,
		map[Coord]int{
			Coord{0, 0}: 4,
			Coord{0, 1}: 4,
			Coord{0, 2}: 4,
			Coord{1, 0}: 4,
			Coord{1, 2}: 4,
			Coord{2, 0}: 4,
			Coord{2, 2}: 4,
		},
	}

	got1 := inputPresent.RotatePresent()

	if maps.Equal(got1.shape, inputPresent.shape) {
		t.Errorf("Failed to rotate present")
	}

	if !maps.Equal(got1.shape, wantPresent1.shape) {
		t.Errorf("Want %v but got %v", wantPresent1.shape, got1.shape)
	}

	/*
		###
		#.#
		#.#

		###
		..#
		###
	*/
	wantPresent2 := Present{
		4,
		map[Coord]int{
			Coord{0, 0}: 4,
			Coord{0, 1}: 4,
			Coord{0, 2}: 4,
			Coord{1, 2}: 4,
			Coord{2, 0}: 4,
			Coord{2, 1}: 4,
			Coord{2, 2}: 4,
		},
	}
	got2 := got1.RotatePresent()
	if !maps.Equal(got2.shape, wantPresent2.shape) {
		t.Errorf("Want %v but got %v", wantPresent2.shape, got2.shape)
	}
}

func TestFlipPresent(t *testing.T) {
	inputPresent := Present{
		4,
		map[Coord]int{
			Coord{0, 0}: 4,
			Coord{0, 1}: 4,
			Coord{0, 2}: 4,
			Coord{1, 0}: 4,
			Coord{2, 0}: 4,
			Coord{2, 1}: 4,
			Coord{2, 2}: 4,
		},
	}

	/*
		###
		#..
		###

		###
		..#
		###
	*/
	wantPresent := Present{
		4,
		map[Coord]int{
			Coord{0, 0}: 4,
			Coord{0, 1}: 4,
			Coord{0, 2}: 4,
			Coord{1, 2}: 4,
			Coord{2, 0}: 4,
			Coord{2, 1}: 4,
			Coord{2, 2}: 4,
		},
	}

	got := inputPresent.FlipPresent()

	if maps.Equal(got.shape, inputPresent.shape) {
		t.Errorf("Failed to flip present")
	}

	if !maps.Equal(got.shape, wantPresent.shape) {
		t.Errorf("Want %v but got %v", wantPresent.shape, got.shape)
	}
}
