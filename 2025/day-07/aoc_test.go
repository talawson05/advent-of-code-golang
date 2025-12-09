package aoc

import (
	"testing"
)

func TestTrackSplitsOnGrid(t *testing.T) {
	fileName := "example_input.txt"
	inputText := ReadAllTextFromFile(fileName)
	got := TrackSplitsOnGrid(inputText)
	want := 21
	if want != got {
		t.Errorf("Wanted %v but got %v", want, got)
	}
}
