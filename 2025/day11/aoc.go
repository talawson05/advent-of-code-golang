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

func SolvePaths(input map[string][]string) int {
	fmt.Println(input)
	return parseTree(input, "you", "out")
}


func Run(fileName string) {
	tree := ParseInput(fileName)
	part1 := SolvePaths(tree)
	fmt.Println("Part1: ", part1)
}