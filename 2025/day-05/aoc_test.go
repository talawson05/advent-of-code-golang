package aoc

import (
	"slices"
	"testing"
)

// Saving for posterity. I liked my fake set thing
/*
func TestRange1to3(t *testing.T) {
	want := []int{1, 2, 3}
	got, err := ExpandRange("1-3")
	if err != nil {
		panic(err)
	}
	if !slices.Equal(want, got) {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func TestNoDupesInSet(t *testing.T) { 
	input := [][]int{{1,2,3}, {3,4,5}}
	want := map[int]bool{
		1: true,
		2: true,
		3: true,
		4: true,
		5: true,
	}
	got := CombineRangesIntoSet(input)
	if !maps.Equal(want, got) {
		t.Errorf("Wanted %v, but got %v", want, got)
	}
}

func TestGapsAreNotPopulatedInSet(t *testing.T) { 
	input := [][]int{{1,2,3}, {5, 6, 7}}
	want := map[int]bool{
		1: true,
		2: true,
		3: true,
		5: true,
		6: true,
		7: true,
	}
	got := CombineRangesIntoSet(input)
	if !maps.Equal(want, got) {
		t.Errorf("Wanted %v, but got %v", want, got)
	}
}

func TestItemIsInSet(t *testing.T) {
	input := [][]int{{1,2,3}, {3,4,5}}
	inputSet := CombineRangesIntoSet(input)
	targetItem := 2
	want := true
	got := ItemIsInSet(inputSet, targetItem)
	if got != want {
		t.Errorf("Wanted %v, but got %v", want, got)
	}
}

func TestItemIsNotInSet(t *testing.T) {
	input := [][]int{{1,2,3}, {3,4,5}}
	inputSet := CombineRangesIntoSet(input)
	targetItem := 9
	want := false
	got := ItemIsInSet(inputSet, targetItem)
	if got != want {
		t.Errorf("Wanted %v, but got %v", want, got)
	}
}

*/

func TestInputParsing(t *testing.T) {
	input := `
	1-3
	6-9

	1
	2
	`
	wantedRanges := []string{"1-3", "6-9"}
	wantedIds := []int{1, 2}
	gotRanges, gotIds := ParseInput(input)

	if (!slices.Equal(wantedRanges, gotRanges)) || !slices.Equal(wantedIds, gotIds) {
		t.Errorf("Wanted \n %v \n with %v IDs \n but got \n %v with %v IDs", wantedRanges, wantedIds, gotRanges, gotIds)
	}
}

func TestNumberInRangeTop(t *testing.T) {
	input := "1-3"
	target := 3
	want := true
	got, err := ItemIsInRange(input, target)
	if err != nil {
		t.Errorf("error thrown: %v", err)
	}

	if got != want {
		t.Errorf("Wanted %v, but got %v", want, got)
	}
}

func TestNumberInRangeMiddle(t *testing.T) {
	input := "1-3"
	target := 2
	want := true
	got, err := ItemIsInRange(input, target)
	if err != nil {
		t.Errorf("error thrown: %v", err)
	}

	if got != want {
		t.Errorf("Wanted %v, but got %v", want, got)
	}
}

func TestNumberInRangeBottom(t *testing.T) {
	input := "1-3"
	target := 1
	want := true
	got, err := ItemIsInRange(input, target)
	if err != nil {
		t.Errorf("error thrown: %v", err)
	}

	if got != want {
		t.Errorf("Wanted %v, but got %v", want, got)
	}
}

func TestNumberNotInRangeLow(t *testing.T) {
	input := "1-3"
	target := 0
	want := false
	got, err := ItemIsInRange(input, target)
	if err != nil {
		t.Errorf("error thrown: %v", err)
	}

	if got != want {
		t.Errorf("Wanted %v, but got %v", want, got)
	}
}

func TestNumberNotInRangeHigh(t *testing.T) {
	input := "1-3"
	target := 9
	want := false
	got, err := ItemIsInRange(input, target)
	if err != nil {
		t.Errorf("error thrown: %v", err)
	}

	if got != want {
		t.Errorf("Wanted %v, but got %v", want, got)
	}
}

func TestCountOfFreshIngredients(t *testing.T) {
	input := `
	3-5
	10-14
	16-20
	12-18

	1
	5
	8
	11
	17
	32
	`
	want := 3
	got := CountOfFreshIngredients(input)
	if got != want {
		t.Errorf("Wanted %v, but got %v", want, got)
	}
}