package aoc

import (
	"reflect"
	"slices"
	"testing"
)

func TestParseStringInputToSlices(t *testing.T) {
	input := `
	123 328  51 64 
	45 64  387 23 
	6 98  215 314
	*   +   *   +  
	`
	wantNumbers := [][]int {
		{123, 328, 51, 64},
		{45, 64, 387, 23},
		{6, 98, 215, 314},
	}
	wantOperators := []string{"*", "+", "*", "+"}

	gotNumbers, gotOperators := ParseInputString(input)

	if !reflect.DeepEqual(wantNumbers, gotNumbers) {
		t.Errorf("Wanted %v, got %v", wantNumbers, gotNumbers)
	}

	if !slices.Equal(gotOperators, wantOperators) {
		t.Errorf("Wanted %v, got %v", wantOperators, gotOperators)
	}
}

func TestCalc(t *testing.T) {
	input := `
	123 328  51 64 
	45 64  387 23 
	6 98  215 314
	*   +   *   +  
	`
	want := 4277556
	got := DoCalc(input)
	if got != want {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}