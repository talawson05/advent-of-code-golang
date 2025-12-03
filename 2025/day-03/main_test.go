package aoc

import (
	"testing"
)

/*
	cd 2025/day-03/
	go test ./...
*/

func TestInner123(t *testing.T) {
	wantIndex, wantValue := 2, "3"
	gotIndex, gotValue := GetIndexAndValueOfBiggestNumberFromRange("123")

	if gotIndex != wantIndex {
		t.Errorf("Wanted index %d, but got %d", wantIndex, gotIndex)
	}
	if gotValue != wantValue{
		t.Errorf("Wanted value %s, but got %s", wantValue, gotValue)
	}
}

func TestInner132(t *testing.T) {
	wantIndex, wantValue := 1, "3"
	gotIndex, gotValue := GetIndexAndValueOfBiggestNumberFromRange("132")

	if gotIndex != wantIndex {
		t.Errorf("Wanted index %d, but got %d", wantIndex, gotIndex)
	}
	if gotValue != wantValue{
		t.Errorf("Wanted value %s, but got %s", wantValue, gotValue)
	}
}

func TestInner321(t *testing.T) {
	wantIndex, wantValue := 0, "3"
	gotIndex, gotValue := GetIndexAndValueOfBiggestNumberFromRange("321")

	if gotIndex != wantIndex {
		t.Errorf("Wanted index %d, but got %d", wantIndex, gotIndex)
	}
	if gotValue != wantValue{
		t.Errorf("Wanted value %s, but got %s", wantValue, gotValue)
	}
}

func TestInner312(t *testing.T) {
	wantIndex, wantValue := 0, "3"
	gotIndex, gotValue := GetIndexAndValueOfBiggestNumberFromRange("312")

	if gotIndex != wantIndex {
		t.Errorf("Wanted index %d, but got %d", wantIndex, gotIndex)
	}
	if gotValue != wantValue{
		t.Errorf("Wanted value %s, but got %s", wantValue, gotValue)
	}
}

func TestWrapper321(t *testing.T) {
	want := "32"
	got := GetBiggestNumberFromRange("321")
	if got != want {
		t.Errorf("Wanted %s, but got %s", want, got)
	}
}

func TestWrapper312(t *testing.T) {
	want := "32"
	got := GetBiggestNumberFromRange("312")
	if got != want {
		t.Errorf("Wanted %s, but got %s", want, got)
	}
}

func TestWrapper231(t *testing.T) {
	want := "31"
	got := GetBiggestNumberFromRange("231")
	if got != want {
		t.Errorf("Wanted %s, but got %s", want, got)
	}
}

func TestWrapper123(t *testing.T) {
	want := "23"
	got := GetBiggestNumberFromRange("123")
	if got != want {
		t.Errorf("Wanted %s, but got %s", want, got)
	}
}

func TestWrapper213(t *testing.T) {
	want := "23"
	got := GetBiggestNumberFromRange("213")
	if got != want {
		t.Errorf("Wanted %s, but got %s", want, got)
	}
}

func TestWrapper987654321111111(t *testing.T) {
	want := "98"
	got := GetBiggestNumberFromRange("987654321111111")
	if got != want {
		t.Errorf("Wanted %s, but got %s", want, got)
	}
}

func TestWrapper811111111111119(t *testing.T) {
	want := "89"
	got := GetBiggestNumberFromRange("811111111111119")
	if got != want {
		t.Errorf("Wanted %s, but got %s", want, got)
	}
}

func TestWrapper234234234234278(t *testing.T) {
	want := "78"
	got := GetBiggestNumberFromRange("234234234234278")
	if got != want {
		t.Errorf("Wanted %s, but got %s", want, got)
	}
}

func TestWrapper818181911112111(t *testing.T) {
	want := "92"
	got := GetBiggestNumberFromRange("818181911112111")
	if got != want {
		t.Errorf("Wanted %s, but got %s", want, got)
	}
}
