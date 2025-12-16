package day10

import (
	"errors"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
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
	start := time.Now()
	limit, parseError := time.ParseDuration("20m")
	if parseError != nil {
		panic("Error parsing duration")
	}
	for {
		// An escape
		elapsed := time.Since(start)
		if elapsed > limit {
			panic("Presses to turn on timed out")
		}

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

const (
	N   = 13
	Eps = 1e-8
)

type Equation struct {
	a [N]float64
	b float64
}

type Variable struct {
	substitution Equation
	independent  bool
	value        int
	upperBound   int
}

func isolate(eq Equation, idx int) (Equation, bool) {
	a := -eq.a[idx]
	if math.Abs(a) < Eps {
		return Equation{}, false
	}

	r := Equation{b: eq.b / a}
	for i := range len(eq.a) {
		if i != idx {
			r.a[i] = eq.a[i] / a
		}
	}
	return r, true
}

func substitute(eq Equation, idx int, expr Equation) Equation {
	r := Equation{}

	a := eq.a[idx]
	eq.a[idx] = 0

	for i := range len(eq.a) {
		r.a[i] = eq.a[i] + a*expr.a[i]
	}
	r.b = eq.b + a*expr.b
	return r
}

func eval(v Variable, vals [N]int) float64 {
	if v.independent {
		return float64(v.value)
	}

	x := v.substitution.b
	for i := range N {
		x += v.substitution.a[i] * float64(vals[i])
	}
	return x
}

func evalRecursive(vars []Variable, free []int, index int) (int, bool) {
	if index == len(free) {
		vals := [N]int{}
		total := 0

		for i := len(vars) - 1; i >= 0; i-- {
			x := eval(vars[i], vals)
			if x < -Eps || math.Abs(x-math.Round(x)) > Eps {
				return 0, false
			}
			vals[i] = int(math.Round(x))
			total += vals[i]
		}

		return total, true
	}

	best, found := math.MaxInt, false
	for x := 0; x <= vars[free[index]].upperBound; x++ {
		vars[free[index]].value = x
		total, ok := evalRecursive(vars, free, index+1)

		if ok {
			found = true
			best = min(best, total)
		}
	}

	if found {
		return best, true
	} else {
		return 0, false
	}
}

func (m *Machine) PressesToMeetJoltage() int {
	vars := make([]Variable, len(m.buttons))

	// Set the initial upper bound for each variable to the maximum possible integer value.
	for i := range vars {
		vars[i].upperBound = math.MaxInt
	}

	// Create equations corresponding to each target joltage.
	equations := make([]Equation, len(m.targetJoltages))
	for i, joltage := range m.targetJoltages {
		equation := Equation{b: float64(-joltage)}
		// Iterate over each button to see if it contributes to this target joltage.
		for j, btn := range m.buttons {
			// Check if the current button's index list contains the current target index.
			if slices.Contains(btn, i) {
				// If so, set the coefficient for this button to 1.
				equation.a[j] = 1
				// Update the upper bound for this variable to the current target joltage.
				vars[j].upperBound = joltage
			}
		}
		// Store the constructed equation.
		equations[i] = equation
	}

	// Attempt to isolate each variable to see if it can be expressed independently.
	for i := range vars {
		vars[i].independent = true
		for _, eq := range equations {
			if expr, ok := isolate(eq, i); ok {
				// If isolation is possible, mark variable as dependent.
				vars[i].independent = false
				// Store the expression that isolates this variable.
				vars[i].substitution = expr
				// Substitute this expression into all other equations to eliminate the variable.
				for j := range equations {
					equations[j] = substitute(equations[j], i, expr)
				}
				break // No need to check further equations once isolated.
			}
		}
	}

	// Collect indices of all independent (free) variables.
	free := make([]int, 0, len(vars))
	for i, v := range vars {
		if v.independent {
			free = append(free, i)
		}
	}

	// Recursively evaluate the minimal number of presses needed, considering free variables.
	best, success := evalRecursive(vars, free, 0)
	if !success {
		panic("Failed to find the best")
	}

	return best
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

	joltageSum := 0
	for _, machine := range machines {
		joltageSum += machine.PressesToMeetJoltage()
	}
	fmt.Println("Part 2:", joltageSum)
}