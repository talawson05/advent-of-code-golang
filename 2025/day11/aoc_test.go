package day11

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	fileName := "example_input.txt"
	got := ParseInput(fileName)
	want := map[string][]string{
		"you" : {"bbb", "ccc"},
		"aaa" : {"you", "hhh"},
		"bbb" : {"ddd", "eee"},
		"ccc" : {"ddd", "eee", "fff"},
		"ddd" : {"ggg",},
		"eee" : {"out",},
		"fff" : {"out",},
		"ggg" : {"out",},
		"hhh" : {"ccc", "fff", "iii"},
		"iii" : {"out",},
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Wanted %v but got %v", want, got)
	}
}

func TestParseTreeSingleChild(t *testing.T) {
	input := map[string][]string{
		"ddd" : {"ggg",},
	}
	got := parseTree(input, "ddd", "ggg")
	want := 1
	if want != got {
		t.Errorf("wanted %v but got %v", want, got)
	}
}

func TestParseTreeTwoChild(t *testing.T) {
	input := map[string][]string{
		"aaa" : {"you", "hhh"},
	}
	got := parseTree(input, "aaa", "you")
	want := 1
	if want != got {
		t.Errorf("wanted %v but got %v", want, got)
	}
}

func TestParseTreeRecursiveOnePath(t *testing.T) {
	input := map[string][]string{
		"you" : {"bbb", "ccc"},
		"bbb": {"ddd", "eee"},
		"fff": {"out"},
	}
	got := parseTree(input, "you", "ddd")
	want := 1
	if want != got {
		t.Errorf("wanted %v but got %v", want, got)
	}
}

func TestParseTreeRecursiveTwoPaths(t *testing.T) {
	input := map[string][]string{
		"you" : {"bbb", "ccc"},
		"bbb": {"ddd", "eee"},
		"ccc": {"ddd", "eee", "fff"},
	}
	got := parseTree(input, "you", "ddd")
	want := 2
	if want != got {
		t.Errorf("wanted %v but got %v", want, got)
	}
}

func TestSolvePaths(t *testing.T) {
	fileName := "example_input.txt"
	tree := ParseInput(fileName)
	got := SolvePaths(tree)
	want := 5
	if want != got {
		t.Errorf("wanted %v but got %v", want, got)
	}
}