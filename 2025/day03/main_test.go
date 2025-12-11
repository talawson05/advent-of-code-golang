package day03

import (
	"slices"
	"testing"
)

/*
	cd 2025/day03/
	go test ./...
*/

func TestWrapper321(t *testing.T) {
	t.Skip("Skipped for part 2")
	want := "32"
	got := GetBiggestNumberFromRange("321")
	if got != want {
		t.Errorf("Wanted %s, but got %s", want, got)
	}
}

func TestWrapper312(t *testing.T) {
	t.Skip("Skipped for part 2")
	want := "32"
	got := GetBiggestNumberFromRange("312")
	if got != want {
		t.Errorf("Wanted %s, but got %s", want, got)
	}
}

func TestWrapper231(t *testing.T) {
	t.Skip("Skipped for part 2")
	want := "31"
	got := GetBiggestNumberFromRange("231")
	if got != want {
		t.Errorf("Wanted %s, but got %s", want, got)
	}
}

func TestWrapper123(t *testing.T) {
	t.Skip("Skipped for part 2")
	want := "23"
	got := GetBiggestNumberFromRange("123")
	if got != want {
		t.Errorf("Wanted %s, but got %s", want, got)
	}
}

func TestWrapper213(t *testing.T) {
	t.Skip("Skipped for part 2")
	want := "23"
	got := GetBiggestNumberFromRange("213")
	if got != want {
		t.Errorf("Wanted %s, but got %s", want, got)
	}
}

func TestWrapper987654321111111(t *testing.T) {
	want := "987654321111"
	got := GetBiggestNumberFromRange("987654321111111")
	if got != want {
		t.Errorf("Wanted %s, but got %s", want, got)
	}
}

func TestWrapper811111111111119(t *testing.T) {
	want := "811111111119"
	got := GetBiggestNumberFromRange("811111111111119")
	if got != want {
		t.Errorf("Wanted %s, but got %s", want, got)
	}
}

func TestWrapper234234234234278(t *testing.T) {
	want := "434234234278"
	got := GetBiggestNumberFromRange("234234234234278")
	if got != want {
		t.Errorf("Wanted %s, but got %s", want, got)
	}
}

func TestWrapper818181911112111(t *testing.T) {
	want := "888911112111"
	got := GetBiggestNumberFromRange("818181911112111")
	if got != want {
		t.Errorf("Wanted %s, but got %s", want, got)
	}
}

func TestConcatIntToString(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	want := "12345"
	got := concatIntsToString(input)
	if got != want {
		t.Errorf("Wanted %v, but got %v", want, got)
	}
}

func TestStringToIntSlice(t *testing.T) {
	want := []int{1, 2, 3, 4, 5}
	input := "12345"
	got := stringToSliceOfInts(input)
	if !slices.Equal(want, got) {
		t.Errorf("Wanted %v, but got %v", want, got)
	}
}
