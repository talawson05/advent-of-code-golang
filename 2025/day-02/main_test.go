package aoc

import (
	"slices"
	"testing"
)

/*
	cd 2025/day-02/
	go test ./...
*/

func Test99IsInvalid(t *testing.T) {
	want := true
	got := IsInvalidId(99)
	if want != got {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func Test22IsInvalid(t *testing.T) {
	want := true
	got := IsInvalidId(22)
	if want != got {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func Test111IsInvalid(t *testing.T) {
	want := true
	got := IsInvalidId(111)
	if want != got {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func Test101IsValid(t *testing.T) {
	want := false
	got := IsInvalidId(101)
	if want != got {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func Test0IsValid(t *testing.T) {
	want := false
	got := IsInvalidId(0)
	if want != got {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func Test1122IsValid(t *testing.T) {
	want := false
	got := IsInvalidId(1122)
	if want != got {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func Test1212IsInvalid(t *testing.T) {
	want := true
	got := IsInvalidId(1212)
	if want != got {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func Test123123123IsInvalid(t *testing.T) {
	want := true
	got := IsInvalidId(123123123)
	if want != got {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func Test1212121212IsInvalid(t *testing.T) {
	want := true
	got := IsInvalidId(1212121212)
	if want != got {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func Test1111111IsInvalid(t *testing.T) {
	want := true
	got := IsInvalidId(1111111)
	if want != got {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func Test2222IsInvalid(t *testing.T) {
	want := true
	got := IsInvalidId(2222)
	if want != got {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func TestRange11to13(t *testing.T) {
	want := []int{11, 12, 13}
	got, err := ExpandRange("11-13")
	if err != nil {
		panic(err)
	}
	if !slices.Equal(want, got) {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func TestSingleNumberInvalid(t *testing.T) {
	want := "invalid input: 42"
	_, got := ExpandRange("42")
	if want != got.Error() {
		t.Errorf("Wanted %v, got %v", want, got)
	}	
}

func TestNoEndNumberInvalid(t *testing.T) {
	//t.Skip("Known issue with split")
	want := "invalid input: 42-"
	_, got := ExpandRange("42-")
	if want != got.Error() {
		t.Errorf("Wanted %v, got %v", want, got)
	}	
}

func TestNoStartingNumberInvalid(t *testing.T) {
	//t.Skip("Known issue with split")
	want := "invalid input: -42"
	_, got := ExpandRange("-42")
	if want != got.Error() {
		t.Errorf("Wanted %v, got %v", want, got)
	}	
}

func TestRange99to101(t *testing.T) {
	want := []int{99, 100, 101}
	got, err := ExpandRange("99-101")
	if err != nil {
		panic(err)
	}
	if !slices.Equal(want, got) {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func TestRange1to2(t *testing.T) {
	want := []int{1, 2}
	got, err := ExpandRange("1-2")
	if err != nil {
		panic(err)
	}
	if !slices.Equal(want, got) {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func TestRange5to5(t *testing.T) {
	want := []int{5}
	got, err := ExpandRange("5-5")
	if err != nil {
		panic(err)
	}
	if !slices.Equal(want, got) {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func TestFindInvalidIdIn11to22(t *testing.T) {
	want := []int{11, 22}
	got, err := ReturnListOfInvalidIdsFromRange("11-22")
	if err != nil {
		panic(err)
	}
	if !slices.Equal(want, got) {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func TestFindInvalidIdIn95to115(t *testing.T) {
	want := []int{99, 111}
	got, err := ReturnListOfInvalidIdsFromRange("95-115")
	if err != nil {
		panic(err)
	}
	if !slices.Equal(want, got) {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func TestFindInvalidIdIn998to1012(t *testing.T) {
	want := []int{999, 1010}
	got, err := ReturnListOfInvalidIdsFromRange("998-1012")
	if err != nil {
		panic(err)
	}
	if !slices.Equal(want, got) {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func TestFindInvalidIdIn1188511880to1188511890(t *testing.T) {
	want := []int{1188511885}
	got, err := ReturnListOfInvalidIdsFromRange("1188511880-1188511890")
	if err != nil {
		panic(err)
	}
	if !slices.Equal(want, got) {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func TestFindInvalidIdIn222220to222224(t *testing.T) {
	want := []int{222222}
	got, err := ReturnListOfInvalidIdsFromRange("222220-222224")
	if err != nil {
		panic(err)
	}
	if !slices.Equal(want, got) {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func TestFindInvalidIdIn1698522to1698528(t *testing.T) {
	want := []int{}
	got, err := ReturnListOfInvalidIdsFromRange("1698522-1698528")
	if err != nil {
		panic(err)
	}
	if !slices.Equal(want, got) {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func TestFindInvalidIdIn446443to446449(t *testing.T) {
	want := []int{446446}
	got, err := ReturnListOfInvalidIdsFromRange("446443-446449")
	if err != nil {
		panic(err)
	}
	if !slices.Equal(want, got) {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func TestFindInvalidIdIn38593856to38593862(t *testing.T) {
	want := []int{38593859}
	got, err := ReturnListOfInvalidIdsFromRange("38593856-38593862")
	if err != nil {
		panic(err)
	}
	if !slices.Equal(want, got) {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func TestSumRange(t *testing.T) {
	want := 6
	got := SumRange([]int{1, 2, 3})
	if want != got {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}
