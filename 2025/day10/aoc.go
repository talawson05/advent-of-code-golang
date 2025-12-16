package day10

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type state bool

func (s *state) toggle() {
	if *s == On {
		*s = Off
	} else {
		*s = On
	}
}

const (
	Off state = false
	On  state = true
)

type Lights []state

func (l Lights) equal(other Lights) bool {
	return slices.Equal(l, other)
}

func (l Lights) press(button Button) Lights {
	result := make(Lights, len(l))
	copy(result, l)
	for _, b := range button {
		result[b].toggle()
	}
	return result
}

type Button []int

type Joltages []int

type Machine struct {
	buttons        []Button
	targetLights   Lights
	targetJoltages Joltages
}

type State struct {
	lights  Lights
	presses int
}

func ParseInput(fileName string) ([]Machine, error) {
	inputText, err := os.ReadFile(fileName)
	if err != nil {
		formatted := fmt.Sprintf("error reading file: %s", err)
		return nil, errors.New(formatted)
	}

	trimmed := strings.TrimSpace(string(inputText))
	lines := strings.Split(trimmed, "\n")

	machines := make([]Machine, 0, len(lines))
	for _, line := range lines {
		// lights
		start := strings.Index(line, "[")
		end := strings.Index(line, "]")
		lights := make(Lights, 0, end-start-1)
		for j := start + 1; j < end; j++ {
			switch line[j] {
			case '.':
				lights = append(lights, Off)
			case '#':
				lights = append(lights, On)
			}
		}

		// buttons
		parts := strings.Split(line, "(")[1:]
		buttons := make([]Button, 0, len(parts))
		for _, part := range parts {
			end := strings.Index(part, ")")
			cleaned := part[:end]
			nums := strings.Split(cleaned, ",")
			button := make(Button, 0, len(nums))
			for _, num := range nums {
				numInt, err := strconv.Atoi(num)
				if err != nil {
					return nil, err
				}
				button = append(button, numInt)
			}
			buttons = append(buttons, button)
		}

		// joltage requirements
		start = strings.Index(line, "{")
		end = strings.Index(line, "}")
		nums := strings.Split(line[start+1:end], ",")
		joltageRequirements := make(Joltages, 0, len(nums))
		for _, num := range nums {
			numInt, err := strconv.Atoi(num)
			if err != nil {
				return nil, err
			}
			joltageRequirements = append(joltageRequirements, numInt)
		}

		machines = append(machines, Machine{
			targetLights:   lights,
			buttons:        buttons,
			targetJoltages: joltageRequirements,
		})

	}
	return machines, nil
}

func (m *Machine) PressesToTurnOn() int {

	startLights := make(Lights, len(m.targetLights))
	queue := []State{{lights: startLights, presses: 0}}

	checked := make(map[string]bool)
	var state State
	for {
		state, queue = queue[0], queue[1:]

		for _, button := range m.buttons {
			newLights := state.lights.press(button)
			hashed := fmt.Sprint(newLights)
			if checked[hashed] {
				// Already checked this combination, step over
				continue
			}
			if newLights.equal(m.targetLights) {
				// Reach correct state, exit early
				return state.presses + 1
			}

			// Save state for next iteration
			newState := State{lights: newLights, presses: state.presses + 1}
			queue = append(queue, newState)
			checked[hashed] = true
		}
	}
}


func Run(fileName string) {
	machines, err := ParseInput(fileName)
	if err != nil {
		panic(err)
	}
	sum := 0
	for _, machine := range machines {
		sum += machine.PressesToTurnOn()
	}
	fmt.Println("Part 1:", sum)
}