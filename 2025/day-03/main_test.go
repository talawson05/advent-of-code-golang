package aoc

import (
	"testing"
)

/*
	cd 2025/day-03/
	go test ./...
*/

func TestInner123(t *testing.T) {
	wantIndex, wantValue := 2, 3
	gotIndex, gotValue := GetIndexAndValueOfBiggestNumberFromRange([]int{1, 2, 3})

	if gotIndex != wantIndex {
		t.Errorf("Wanted index %d, but got %d", wantIndex, gotIndex)
	}
	if gotValue != wantValue{
		t.Errorf("Wanted value %d, but got %d", wantValue, gotValue)
	}
}

func TestInner132(t *testing.T) {
	wantIndex, wantValue := 1, 3
	gotIndex, gotValue := GetIndexAndValueOfBiggestNumberFromRange([]int{1, 3, 2})

	if gotIndex != wantIndex {
		t.Errorf("Wanted index %d, but got %d", wantIndex, gotIndex)
	}
	if gotValue != wantValue{
		t.Errorf("Wanted value %d, but got %d", wantValue, gotValue)
	}
}

func TestInner321(t *testing.T) {
	wantIndex, wantValue := 0, 3
	gotIndex, gotValue := GetIndexAndValueOfBiggestNumberFromRange([]int{3,2,1})

	if gotIndex != wantIndex {
		t.Errorf("Wanted index %d, but got %d", wantIndex, gotIndex)
	}
	if gotValue != wantValue{
		t.Errorf("Wanted value %d, but got %d", wantValue, gotValue)
	}
}

func TestInner312(t *testing.T) {
	wantIndex, wantValue := 0, 3
	gotIndex, gotValue := GetIndexAndValueOfBiggestNumberFromRange([]int{3,1,2})

	if gotIndex != wantIndex {
		t.Errorf("Wanted index %d, but got %d", wantIndex, gotIndex)
	}
	if gotValue != wantValue{
		t.Errorf("Wanted value %d, but got %d", wantValue, gotValue)
	}
}

func TestWrapper321(t *testing.T) {
	want := 32
	got := GetBiggestNumberFromRange([]int{3,2,1})
	if got != want {
		t.Errorf("Wanted %d, but got %d", want, got)
	}
}

func TestWrapper312(t *testing.T) {
	want := 32
	got := GetBiggestNumberFromRange([]int{3,1,2})
	if got != want {
		t.Errorf("Wanted %d, but got %d", want, got)
	}
}

func TestWrapper231(t *testing.T) {
	want := 31
	got := GetBiggestNumberFromRange([]int{2,3,1})
	if got != want {
		t.Errorf("Wanted %d, but got %d", want, got)
	}
}

func TestWrapper123(t *testing.T) {
	want := 23
	got := GetBiggestNumberFromRange([]int{1,2,3})
	if got != want {
		t.Errorf("Wanted %d, but got %d", want, got)
	}
}

func TestWrapper213(t *testing.T) {
	want := 23
	got := GetBiggestNumberFromRange([]int{2,1,3})
	if got != want {
		t.Errorf("Wanted %d, but got %d", want, got)
	}
}
