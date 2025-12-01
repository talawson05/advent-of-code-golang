package aoc

import (
	"fmt"
	"testing"
	"os"
)

func TestSimpleRightTurn (t *testing.T) {
	want := 12
	got, err := DoRotation(10, "R", 2)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if want != got {
		t.Errorf("Wanted %d, got %d", want, got)
	}
}

func TestSimpleLeftTurn (t *testing.T) {
	want := 8
	got, err := DoRotation(10, "L", 2)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if want != got {
		t.Errorf("Wanted %d, got %d", want, got)
	}
}

func TestCompletingTheCircleRight(t *testing.T) {
	want := 0
	got, err := DoRotation(99, "R", 1)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if want != got {
		t.Errorf("Wanted %d, got %d", want, got)
	}
}

func TestCompletingTheCircleLeft(t *testing.T) {
	want := 99
	got, err := DoRotation(0, "L", 1)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if want != got {
		t.Errorf("Wanted %d, got %d", want, got)
	}
}

func TestLargeRightTurn(t *testing.T) {
	want := 30
	got, err := DoRotation(10, "R", 120)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if want != got {
		t.Errorf("Wanted %d, got %d", want, got)
	}
}

func TestLargeLeftTurn(t *testing.T) {
	want := 65
	got, err := DoRotation(10, "L", 145)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if want != got {
		t.Errorf("Wanted %d, got %d", want, got)
	}
}


func TestGoingNowhereRight(t *testing.T) {
	want := 42
	got, err := DoRotation(42, "R", 0)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if want != got {
		t.Errorf("Wanted %d, got %d", want, got)
	}
}

func TestGoingNowhereLeft(t *testing.T) {
	want := 42
	got, err := DoRotation(42, "L", 0)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if want != got {
		t.Errorf("Wanted %d, got %d", want, got)
	}
}


func TestUnexpectedDirection(t *testing.T) {
	expectError := "unexpected direction"
	_, actualError := DoRotation(33, "U", 1)

	if (actualError == nil) || (actualError.Error() != expectError) {
		t.Errorf("Wanted %s, got %v", expectError, actualError)
	}
}
