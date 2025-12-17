package day11

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ParseInput(fileName string) map[string][]string {
	returnValue := map[string][]string{}
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close() //nolint:errcheck

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()
		values := strings.Split(currentLine, ":")
		returnValue[values[0]] = strings.Fields(values[1])
	}

	return returnValue
}

func parseTree(tree map[string][]string, start, end string) int {
	count := 0
	currentNode := tree[start]
	for _, child := range currentNode {
		if child == end {
			// winner
			return 1
		}
		// keep going
		count += parseTree(tree, child, end)
	}
	return count
}

func parseTreePart2(tree map[string][]string, start, end string, passedDac, passedFft bool, cache map[string]int) int {
	count := 0
	currentNode := tree[start]
	if cacheValue, ok := cache[fmt.Sprintf("%v:%v:%v:%v", start, currentNode, passedDac, passedFft)]; ok {
		return cacheValue
	}

	if start == end {
		if passedDac && passedFft {
			// winner
			return 1
		}
		// reached end but didn't pass both dac and fft
		return 0
	}

	switch start {
	case "dac":
		passedDac = true
	case "fft":
		passedFft = true
	}

	for _, child := range currentNode {
		count += parseTreePart2(tree, child, end, passedDac, passedFft, cache)
	}
	cache[fmt.Sprintf("%v:%v:%v:%v", start, currentNode, passedDac, passedFft)] = count
	return count
}

func SolvePaths(input map[string][]string) int {
	return parseTree(input, "you", "out")
}

func SolvePathsPart2(input map[string][]string) int {
	passedDac, passedFft := false, false
	cache := map[string]int{}
	return parseTreePart2(input, "svr", "out", passedDac, passedFft, cache)
}

func Run(fileName string) {
	tree := ParseInput(fileName)
	part1 := SolvePaths(tree)
	fmt.Println("Part1: ", part1)
	part2 := SolvePathsPart2(tree)
	fmt.Println("Part2: ", part2)
}
