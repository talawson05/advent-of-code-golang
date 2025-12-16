package day10

import (
	"slices"
	"testing"
)

func TestParseInput(t *testing.T) {
	inputFile := "example_input.txt"
	got, err := ParseInput(inputFile)
	if err != nil {
		t.Errorf("Got error: %s", err)
	}

	if len(got) != 3 {
		t.Errorf("Wanted 3 machines, but got %v", len(got))
	}

	machine0Lights := len(got[0].targetLights)
	if machine0Lights != 4 {
		t.Errorf("Wanted 4 lights, but got %v", machine0Lights)
	}

	machine1Buttons := len(got[1].buttons)
	if machine1Buttons != 5 {
		t.Errorf("Wanted 5 buttons, but got %v", machine1Buttons)
	}
}

func TestLightToggleState(t *testing.T) {
	state := State{lights: Lights{Off, Off, Off, Off}}
	buttons := []Button{{3}, {1,3}, {2}, {2,3}, {0,2}, {0,1},}

	state.lights.press(buttons[0]) // 3
	expectedAfterPress0 := Lights{Off, Off, Off, On}
	if slices.Equal(state.lights, expectedAfterPress0) {
		t.Errorf("Expected %v but got %v", expectedAfterPress0, state.lights)
	}

	state.lights.press(buttons[1]) // 1,3
	expectedAfterPress1 := Lights{Off, On, Off, Off}
	if slices.Equal(state.lights, expectedAfterPress1) {
		t.Errorf("Expected %v but got %v", expectedAfterPress1, state.lights)
	}

	state.lights.press(buttons[2]) // 2
	expectedAfterPress2 := Lights{Off, On, On, Off}
	if slices.Equal(state.lights, expectedAfterPress2) {
		t.Errorf("Expected %v but got %v", expectedAfterPress2, state.lights)
	}

	state.lights.press(buttons[3]) // 2,3
	expectedAfterPress3 := Lights{Off, On, Off, On}
	if slices.Equal(state.lights, expectedAfterPress3) {
		t.Errorf("Expected %v but got %v", expectedAfterPress3, state.lights)
	}
}

func TestExampleInput0(t *testing.T) {
	inputFile := "example_input.txt"
	machines, err := ParseInput(inputFile)
	if err != nil {
		t.Errorf("Got error: %s", err)
	}

	got := machines[0].PressesToTurnOn()
	want := 2

	if want != got {
		t.Errorf("Wanted %v but got %v", want, got)
	}
}

func TestExampleInput1(t *testing.T) {
	inputFile := "example_input.txt"
	machines, err := ParseInput(inputFile)
	if err != nil {
		t.Errorf("Got error: %s", err)
	}

	got := machines[1].PressesToTurnOn()
	want := 3

	if want != got {
		t.Errorf("Wanted %v but got %v", want, got)
	}
}

func TestExampleInput2(t *testing.T) {
	inputFile := "example_input.txt"
	machines, err := ParseInput(inputFile)
	if err != nil {
		t.Errorf("Got error: %s", err)
	}

	got := machines[2].PressesToTurnOn()
	want := 2

	if want != got {
		t.Errorf("Wanted %v but got %v", want, got)
	}
}

