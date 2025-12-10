package aoc

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Run(fileName string, targetNumberOfConnections int) int {
	input := ParseInput(fileName)
	pairs := CalculateDistanceForAllJunctionBoxes(input)
	SortJunctionBoxPairsByDistance(pairs)

	circuits := []Circuit{}
	for i := range targetNumberOfConnections {
		currentPair := pairs[i]
		// fmt.Println(currentPair)
		returnedCircuits, returnedPairs, returnedPair := ConnectPair(circuits, pairs, currentPair)
		// fmt.Printf("Returned circuit was %v\n", returnedCircuits)
		// fmt.Printf("Returned paif was %v\n", returnedPair)
		// pairs = UpdatePairsAfterConnection(returnedPairs, returnedPair)
		// fmt.Println(pairs)

		circuits = returnedCircuits
		pairs = returnedPairs
		pairs[i] = returnedPair
		// fmt.Printf("main circuits now %v\n", circuits)
	}

	// fmt.Println(circuits)

	// Want the largest 3 circuits
	largeCircuits := GetXCircuitsBySize(circuits, 3)
	// fmt.Println("Now the largest ones")
	// fmt.Println(largeCircuits)
	product := MultiplyNumberOfJunctionBoxesInCircuits(largeCircuits)
	fmt.Println(product)
	return product
}

// circuitId of 0 means not in a circuit
type JunctionBox struct {
	x, y, z, circuitId int
}

type Pair struct {
	jb1, jb2 JunctionBox
	distance int
}

type Circuit struct {
	id  int
	jbs []JunctionBox
}

func ParseInput(fileName string) []JunctionBox {
	returnValue := []JunctionBox{}

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()
		values := strings.Split(currentLine, ",")
		// fmt.Println(values)

		xPos, _ := strconv.Atoi(values[0])
		yPos, _ := strconv.Atoi(values[1])
		zPos, _ := strconv.Atoi(values[2])
		jb := JunctionBox{xPos, yPos, zPos, -1}
		// fmt.Println(jb)
		returnValue = append(returnValue, jb)
	}

	return returnValue
}

func CalcDistanceBetweenTwoJunctionBoxes(jb1, jb2 JunctionBox) int {
	dX := jb1.x - jb2.x
	dY := jb1.y - jb2.y
	dZ := jb1.z - jb2.z
	distance := dX*dX + dY*dY + dZ*dZ
	// fmt.Println(distance)
	return distance
}

func CalculateDistanceForAllJunctionBoxes(input []JunctionBox) []Pair {
	returnValue := []Pair{}

	// nested loop to check every item against every other item
	length := len(input)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			firstJunctionBox, secondJunctionBox := input[i], input[j]
			// fmt.Printf("Working with %v and %v\n", firstJunctionBox, secondJunctionBox)
			dist := CalcDistanceBetweenTwoJunctionBoxes(firstJunctionBox, secondJunctionBox)
			currentPair := Pair{firstJunctionBox, secondJunctionBox, dist}
			returnValue = append(returnValue, currentPair)
		}
	}

	return returnValue
}

// No return type as it mutates orginal slice
func SortJunctionBoxPairsByDistance(input []Pair) {
	sort.Slice(input, func(a, b int) bool { return input[a].distance < input[b].distance })
}

func GetNextCircuitId(circuits []Circuit) int {
	return len(circuits)
}

func ConnectPair(circuits []Circuit, pairs []Pair, pairToConnect Pair) ([]Circuit, []Pair, Pair) {

	firstJbCircuit := pairToConnect.jb1.circuitId
	secondJbCircuit := pairToConnect.jb2.circuitId

	// Not connected
	if firstJbCircuit == -1 && secondJbCircuit == -1 {
		circuitId := GetNextCircuitId(circuits)

		pairToConnect.jb1.circuitId = circuitId
		pairToConnect.jb2.circuitId = circuitId

		pairs = UpdatePairsAfterConnection(pairs, pairToConnect)

		newCircuit := Circuit{
			circuitId,
			[]JunctionBox{pairToConnect.jb1, pairToConnect.jb2},
		}
		circuits = append(circuits, newCircuit)

		// fmt.Println("Connected in new circuit")
		// fmt.Println(pair)

	} else if firstJbCircuit != -1 && secondJbCircuit == -1 {
		// join first network
		pairToConnect.jb2.circuitId = firstJbCircuit
		pairs = UpdatePairsAfterConnection(pairs, pairToConnect)
		circuits[firstJbCircuit].jbs = append(circuits[firstJbCircuit].jbs, pairToConnect.jb2)
	} else if firstJbCircuit == -1 && secondJbCircuit != -1 {
		// Join second network
		pairToConnect.jb1.circuitId = secondJbCircuit
		pairs = UpdatePairsAfterConnection(pairs, pairToConnect)
		circuits[secondJbCircuit].jbs = append(circuits[secondJbCircuit].jbs, pairToConnect.jb1)

	} else if firstJbCircuit == secondJbCircuit {
		// Both in same network, do nothing
		return circuits, pairs, pairToConnect
	} else if firstJbCircuit != secondJbCircuit {
		// do diff circuits, combine?

		// merge all JunctionBoxes from circuit 2 into circuit 1
		pairToConnect.jb2.circuitId = firstJbCircuit
		circuits[firstJbCircuit].jbs = append(circuits[firstJbCircuit].jbs, circuits[secondJbCircuit].jbs...)
		circuits = append(circuits[:secondJbCircuit], circuits[secondJbCircuit:]...) // remove
		pairs = UpdateAllPairsFromOldCircuitToNewCircuit(pairs, secondJbCircuit, firstJbCircuit)
	}

	return circuits, pairs, pairToConnect
}

func UpdatePairsAfterConnection(listOfPairs []Pair, udpatedPair Pair) []Pair {

	listOfPairs = UpdateJunctionBoxInPairs(listOfPairs, udpatedPair.jb1)
	listOfPairs = UpdateJunctionBoxInPairs(listOfPairs, udpatedPair.jb2)

	return listOfPairs
}

func UpdateAllPairsFromOldCircuitToNewCircuit(listOfPairs []Pair, oldCircuitId, newCircuitId int) []Pair {
	for p := range listOfPairs {
		if listOfPairs[p].jb1.circuitId == oldCircuitId {
			tempJob := listOfPairs[p].jb1
			tempJob.circuitId = newCircuitId
			listOfPairs = UpdateJunctionBoxInPairs(listOfPairs, tempJob)
		}
		if listOfPairs[p].jb2.circuitId == oldCircuitId {
			tempJob := listOfPairs[p].jb2
			tempJob.circuitId = newCircuitId
			listOfPairs = UpdateJunctionBoxInPairs(listOfPairs, tempJob)
		}
	}
	return listOfPairs
}

func UpdateJunctionBoxInPairs(listOfPairs []Pair, updatedJunctionBox JunctionBox) []Pair {
	for index, currentPair := range listOfPairs {
		if currentPair.jb1.x == updatedJunctionBox.x &&
			currentPair.jb1.y == updatedJunctionBox.y &&
			currentPair.jb1.z == updatedJunctionBox.z {
			listOfPairs[index].jb1 = updatedJunctionBox
		} else if currentPair.jb2.x == updatedJunctionBox.x &&
			currentPair.jb2.y == updatedJunctionBox.y &&
			currentPair.jb2.z == updatedJunctionBox.z {
			listOfPairs[index].jb2 = updatedJunctionBox
		}
	}
	return listOfPairs
}

func GetXCircuitsBySize(inputCircuits []Circuit, xAmount int) []Circuit {
	sort.Slice(inputCircuits, func(a, b int) bool {
		return len(inputCircuits[a].jbs) > len(inputCircuits[b].jbs)
	})
	return inputCircuits[:xAmount]
}

func MultiplyNumberOfJunctionBoxesInCircuits(inputCircuits []Circuit) int {
	returnValue := 1

	for _, circuit := range inputCircuits {
		numberOfJenctionBoxes := len(circuit.jbs)
		returnValue *= numberOfJenctionBoxes
	}

	return returnValue
}
