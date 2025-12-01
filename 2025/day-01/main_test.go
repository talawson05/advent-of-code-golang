package aoc

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"testing"
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

func TestExampleInput(t *testing.T) {
	want := []string{"L68", "L30", "R48", "L5", "R60","L55", "L1","L99", "R14", "L82"}
	got := []string{}

	file, err := os.Open("example_input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()
		got = append(got, currentLine)
	}

	if !slices.Equal(want, got) {
		t.Errorf("Wanted %v, got %v", want, got)
	}

	// rename for readability
	steps := got

	// assert step by step
	startingPosition := 50
	step0 := steps[0] // L68
	step0Clicks, step0ParseErr := strconv.Atoi(step0[1:])
	if step0ParseErr != nil {
        panic(step0ParseErr)
    }

	positionAfterStep0, step0Err := DoRotation(startingPosition, step0[:1], step0Clicks)
	if step0Err != nil {
		panic(step0Err)
	}
	if positionAfterStep0 != 82 {
		t.Errorf("Wanted 82, got %d", positionAfterStep0)
	}

	step1 := steps[1] // L30
	step1Clicks, step1ParseErr := strconv.Atoi(step1[1:])
	if step1ParseErr != nil {
        panic(step1ParseErr)
    }
	positionAfterStep1, step1Err := DoRotation(positionAfterStep0, step1[:1], step1Clicks)
	if step1Err != nil {
		panic(step1Err)
	}
	if positionAfterStep1 != 52 {
		t.Errorf("Wanted 52, got %d", positionAfterStep1)
	}

	step2 := steps[2] // R48
	step2Clicks, step2ParseErr := strconv.Atoi(step2[1:])
	if step2ParseErr != nil {
        panic(step2ParseErr)
    }
	positionAfterStep2, step2Err := DoRotation(positionAfterStep1, step2[:1], step2Clicks)
	if step2Err != nil {
		panic(step2Err)
	}
	if positionAfterStep2 != 0 {
		t.Errorf("Wanted 0, got %d", positionAfterStep2)
	}

	step3 := steps[3] // L5
	step3Clicks, step3ParseErr := strconv.Atoi(step3[1:])
	if step3ParseErr != nil {
        panic(step3ParseErr)
    }
	positionAfterStep3, step3Err := DoRotation(positionAfterStep2, step3[:1], step3Clicks)
	if step3Err != nil {
		panic(step3Err)
	}
	if positionAfterStep3 != 95 {
		t.Errorf("Wanted 95, got %d", positionAfterStep3)
	}

	step4 := steps[4] // R60
	step4Clicks, step4ParseErr := strconv.Atoi(step4[1:])
	if step4ParseErr != nil {
        panic(step4ParseErr)
    }
	positionAfterStep4, step4Err := DoRotation(positionAfterStep3, step4[:1], step4Clicks)
	if step4Err != nil {
		panic(step4Err)
	}
	if positionAfterStep4 != 55 {
		t.Errorf("Wanted 55, got %d", positionAfterStep4)
	}




}

