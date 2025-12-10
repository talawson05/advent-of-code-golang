package aoc

import (
	"fmt"
	"slices"
	"testing"
)

func TestParseInput(t *testing.T) {
	inputFile := "example_input.txt"
	got := ParseInput(inputFile)
	want := []JunctionBox{
		{162, 817, 812, -1},
		{57, 618, 57, -1},
		{906, 360, 560, -1},
		{592, 479, 940, -1},
		{352, 342, 300, -1},
		{466, 668, 158, -1},
		{542, 29, 236, -1},
		{431, 825, 988, -1},
		{739, 650, 466, -1},
		{52, 470, 668, -1},
		{216, 146, 977, -1},
		{819, 987, 18, -1},
		{117, 168, 530, -1},
		{805, 96, 715, -1},
		{346, 949, 466, -1},
		{970, 615, 88, -1},
		{941, 993, 340, -1},
		{862, 61, 35, -1},
		{984, 92, 344, -1},
		{425, 690, 689, -1},
	}

	if !slices.Equal(want, got) {
		t.Errorf("Wanted %v but got %v", want, got)
	}
}

func TestDistanceCalc(t *testing.T) {
	jb1 := JunctionBox{162, 817, 812, -1}
	jb2 := JunctionBox{425, 690, 689, -1}
	want := 100427
	got := CalcDistanceBetweenTwoJunctionBoxes(jb1, jb2)
	if want != got {
		t.Errorf("Wanted %v but got %v", want, got)
	}
}

func TestListOfDistancesCalculation(t *testing.T) {
	inputFile := "example_input.txt"
	junctionBoxes := ParseInput(inputFile)
	got := CalculateDistanceForAllJunctionBoxes(junctionBoxes)
	// too many permutations to create exact expected output
	// use contains to spot check
	jb1 := JunctionBox{162, 817, 812, -1}
	jb2 := JunctionBox{425, 690, 689, -1}
	expectedPair := Pair{jb1: jb1, jb2: jb2, distance: 100427}
	// Cover duplicate edge case
	unexpectedPair := Pair{jb2, jb1, 100427}

	if !slices.Contains(got, expectedPair) {
		t.Errorf("Expected to find %v in %v", expectedPair, got)
	}

	if slices.Contains(got, unexpectedPair) {
		t.Errorf("Unexpected value %v in %v", unexpectedPair, got)
	}
}

func TestSortByDistance(t *testing.T) {
	inputFile := "example_input.txt"
	junctionBoxes := ParseInput(inputFile)
	listOfPairs := CalculateDistanceForAllJunctionBoxes(junctionBoxes)
	// fmt.Println(listOfPairs)
	SortJunctionBoxPairsByDistance(listOfPairs)
	// fmt.Println("after")
	// fmt.Println(listOfPairs)
	got := listOfPairs[:2]
	firstPair := Pair{JunctionBox{162, 817, 812, -1}, JunctionBox{425, 690, 689, -1}, 100427}
	secondPair := Pair{JunctionBox{162, 817, 812, -1}, JunctionBox{431, 825, 988, -1}, 103401}
	expectedValue := []Pair{firstPair, secondPair}
	if !slices.Equal(got, expectedValue) {
		t.Errorf("Expected %v but got %v", expectedValue, got)
	}
}

func TestGetNextCircuitIdEmpty(t *testing.T) {
	listOfCircuits := []Circuit{}
	want := 0
	got := GetNextCircuitId(listOfCircuits)
	if want != got {
		t.Errorf("Wanted %v but got %v", want, got)
	}
}

func TestGetNextCircuitIdPopulated(t *testing.T) {
	listOfCircuits := []Circuit{
		{
			1,
			[]JunctionBox{JunctionBox{162, 817, 812, 0}, JunctionBox{425, 690, 689, 0}},
		},
	}
	want := 1
	got := GetNextCircuitId(listOfCircuits)
	if want != got {
		t.Errorf("Wanted %v but got %v", want, got)
	}
}

func TestConnectPairNoCircuit(t *testing.T) {
	listOfCircuits := []Circuit{}
	listOfPairs := []Pair{
		Pair{JunctionBox{162, 817, 812, -1}, JunctionBox{425, 690, 689, -1}, 100427},
	}
	inputPair := listOfPairs[0]
	listOfCircuits, _, inputPair = ConnectPair(listOfCircuits, listOfPairs, inputPair)

	if len(listOfCircuits) != 1 {
		t.Errorf("Expected %v in %v", 1, listOfCircuits)
	}

	if inputPair.jb1.circuitId != 0 {
		t.Errorf("Expected %v in %v", 0, inputPair.jb1.circuitId)
	}

	if inputPair.jb2.circuitId != 0 {
		t.Errorf("Expected %v in %v", 0, inputPair.jb2.circuitId)
	}
}

func TestConnectPairOneConnected(t *testing.T) {
	listOfCircuits := []Circuit{
		{
			0,
			[]JunctionBox{
				JunctionBox{162, 817, 812, 0},
				JunctionBox{425, 690, 689, 0},
			},
		},
	}
	listOfPairs := []Pair{
		Pair{JunctionBox{162, 817, 812, 0}, JunctionBox{425, 690, 689, 0}, 100427},
		Pair{JunctionBox{162, 817, 812, 0}, JunctionBox{431, 825, 988, -1}, 103401},
	}
	inputPair := Pair{JunctionBox{162, 817, 812, 0}, JunctionBox{431, 825, 988, -1}, 103401}
	listOfCircuits, _, inputPair = ConnectPair(listOfCircuits, listOfPairs, inputPair)

	if len(listOfCircuits) != 1 {
		t.Errorf("Expected %v in %v", 1, listOfCircuits)
	}

	if inputPair.jb1.circuitId != 0 {
		t.Errorf("Expected %v in %v", 0, inputPair.jb1.circuitId)
	}

	if inputPair.jb2.circuitId != 0 {
		t.Errorf("Expected %v in %v", 0, inputPair.jb2.circuitId)
	}
}

func TestListOfPairsIsUpdatedAfterConnecting(t *testing.T) {

	listOfCircuits := []Circuit{}
	listofPairs := []Pair{
		Pair{JunctionBox{162, 817, 812, -1}, JunctionBox{425, 690, 689, -1}, 100427},
		Pair{JunctionBox{162, 817, 812, -1}, JunctionBox{431, 825, 988, -1}, 103401},
	}
	inputPair := listofPairs[0]
	_, _, connectedPair := ConnectPair(listOfCircuits, listofPairs, inputPair)
	listofPairs = UpdatePairsAfterConnection(listofPairs, connectedPair)

	updatedPair := listofPairs[1]

	// JB1 should be updated
	if updatedPair.jb1.circuitId != 0 {
		t.Errorf("Expected %v in %v", 0, updatedPair.jb1.circuitId)
	}

	// JB2 should not be updated
	if updatedPair.jb2.circuitId != -1 {
		t.Errorf("Expected %v in %v", 0, updatedPair.jb2.circuitId)
	}

}

func TestUpdateAllPairsFromOldCircuitToNewCircuit(t *testing.T) {
	listOfPairs := []Pair{
		Pair{JunctionBox{162, 817, 812, 0}, JunctionBox{425, 690, 689, 0}, 100427},
		Pair{JunctionBox{162, 817, 812, 0}, JunctionBox{431, 825, 988, 1}, 103401}, // join
		Pair{JunctionBox{999, 999, 999, 1}, JunctionBox{666, 666, 666, 1}, 42},     // to be updated
		Pair{JunctionBox{222, 222, 222, 2}, JunctionBox{22, 22, 22, 2}, 42},        // to not be updated
	}
	returnedList := UpdateAllPairsFromOldCircuitToNewCircuit(listOfPairs, 1, 0)

	if returnedList[1].jb2.circuitId != 0 {
		t.Errorf("Expected JB circuit id to be updated: %v", returnedList[1].jb2.circuitId)
	}
	if returnedList[2].jb1.circuitId != 0 {
		t.Errorf("Expected JB circuit id to be updated: %v", returnedList[2].jb1.circuitId)
	}
	if returnedList[2].jb2.circuitId != 0 {
		t.Errorf("Expected JB circuit id to be updated: %v", returnedList[2].jb2.circuitId)
	}
	if returnedList[3].jb1.circuitId != 2 {
		t.Errorf("Expected JB circuit id to not be updated: %v", returnedList[3].jb1.circuitId)
	}
	if returnedList[3].jb2.circuitId != 2 {
		t.Errorf("Expected JB circuit id to not be updated: %v", returnedList[3].jb2.circuitId)
	}
}

func TestConnectionBothSameCircuit(t *testing.T) {
	listOfCircuits := []Circuit{
		{
			0,
			[]JunctionBox{
				JunctionBox{162, 817, 812, 0},
				JunctionBox{425, 690, 689, 0},
				JunctionBox{431, 825, 988, 0},
			},
		},
	}
	listOfPairs := []Pair{
		Pair{JunctionBox{431, 825, 988, 0}, JunctionBox{425, 690, 689, 0}, 42},
	}
	inputPair := listOfPairs[0]
	listOfCircuits, _, inputPair = ConnectPair(listOfCircuits, listOfPairs, inputPair)

	if len(listOfCircuits) != 1 {
		t.Errorf("Expected %v in %v", 1, listOfCircuits)
	}

	if inputPair.jb1.circuitId != 0 {
		t.Errorf("Expected %v in %v", 0, inputPair.jb1.circuitId)
	}

	if inputPair.jb2.circuitId != 0 {
		t.Errorf("Expected %v in %v", 0, inputPair.jb2.circuitId)
	}
}

func TestConnectTwoNetworks(t *testing.T) {
	listOfCircuits := []Circuit{
		{
			0,
			[]JunctionBox{
				JunctionBox{162, 817, 812, 0},
				JunctionBox{425, 690, 689, 0},
				JunctionBox{431, 825, 988, 0},
			},
		},
		{
			1,
			[]JunctionBox{
				JunctionBox{906, 360, 560, 1},
				JunctionBox{805, 96, 715, 1},
				JunctionBox{999, 999, 999, 1},
			},
		},
	}
	listOfPairs := []Pair{
		{JunctionBox{162, 817, 812, 0}, JunctionBox{425, 690, 689, 0}, 42},
		{JunctionBox{162, 817, 812, 0}, JunctionBox{431, 825, 988, 0}, 42},
		{JunctionBox{906, 360, 560, 1}, JunctionBox{805, 96, 715, 1}, 42},
		{JunctionBox{805, 96, 715, 1}, JunctionBox{999, 999, 999, 1}, 42},
		{JunctionBox{431, 825, 988, 0}, JunctionBox{999, 999, 999, 1}, 42}, // this one
	}
	inputPair := listOfPairs[4]
	returnedListOfCircuits, returnedListOfPairs, returnedPair := ConnectPair(listOfCircuits, listOfPairs, inputPair)

	want := Pair{JunctionBox{431, 825, 988, 0}, JunctionBox{999, 999, 999, 0}, 42}
	if returnedPair != want {
		t.Errorf("Expected %v in %v", want, returnedPair)
	}

	if len(returnedListOfCircuits) != 1 {
		t.Errorf("Expected %v in %v", 1, listOfCircuits)
	}

	if returnedListOfPairs[2].jb1.circuitId != 0 {
		t.Errorf("Expected %v in %v", 0, returnedListOfPairs[2].jb1.circuitId)
	}

	if returnedListOfPairs[2].jb2.circuitId != 0 {
		t.Errorf("Expected %v in %v", 0, returnedListOfPairs[2].jb2.circuitId)
	}
}

func TestGetXNumberOfCircuitsBySize(t *testing.T) {
	listOfCircuits := []Circuit{
		{
			1,
			[]JunctionBox{
				JunctionBox{906, 360, 560, 1},
				JunctionBox{805, 96, 715, 1},
				JunctionBox{999, 999, 999, 1},
			},
		},
		{
			2,
			[]JunctionBox{
				JunctionBox{222, 222, 222, 2},
				JunctionBox{22, 22, 22, 2},
			},
		},
		{
			0,
			[]JunctionBox{
				JunctionBox{162, 817, 812, 0},
				JunctionBox{425, 690, 689, 0},
				JunctionBox{431, 825, 988, 0},
			},
		},
	}

	returnedCircuits := GetXCircuitsBySize(listOfCircuits, 2)
	wantLength := 2

	if len(returnedCircuits) != wantLength {
		t.Errorf("Wanted %v but got %v", wantLength, len(returnedCircuits))
	}

	fmt.Println(returnedCircuits)

}

func TestMultiplyNumberOfJunctionBoxesInCircuits(t *testing.T) {
	listOfCircuits := []Circuit{
		{
			0,
			[]JunctionBox{
				JunctionBox{162, 817, 812, 0},
				JunctionBox{425, 690, 689, 0},
				JunctionBox{431, 825, 988, 0},
			},
		},
		{
			1,
			[]JunctionBox{
				JunctionBox{906, 360, 560, 1},
				JunctionBox{805, 96, 715, 1},
				JunctionBox{999, 999, 999, 1},
			},
		},
		{
			2,
			[]JunctionBox{
				JunctionBox{222, 222, 222, 2},
				JunctionBox{22, 22, 22, 2},
			},
		},
	}

	got := MultiplyNumberOfJunctionBoxesInCircuits(listOfCircuits)

	// 3 * 3 * 2 = 18
	want := 18
	if want != got {
		t.Errorf("Wanted %v but got %v", want, got)
	}

}

func TestExampleInputForPart1(t *testing.T) {
	fileName := "example_input.txt"
	got := Run(fileName, 10)
	want := 40
	if want != got {
		t.Errorf("Wanted %v but got %v", want, got)
	}
}
