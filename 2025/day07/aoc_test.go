package day07

import (
	"testing"
)

func TestTrackSplitsOnGrid(t *testing.T) {
	fileName := "example_input.txt"
	inputText := ReadAllTextFromFile(fileName)
	got, _ := TrackSplitsOnGrid(inputText)
	want := 21
	if want != got {
		t.Errorf("Wanted %v but got %v", want, got)
	}
}

func TestTimelineSplitsOnGrid(t *testing.T) {
	fileName := "example_input.txt"
	inputText := ReadAllTextFromFile(fileName)
	_, got := TrackSplitsOnGrid(inputText)
	want := 40
	if want != got {
		t.Errorf("Wanted %v but got %v", want, got)
	}
}
