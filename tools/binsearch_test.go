package tools

import "testing"

func TestBinarySearch(t *testing.T) {
	myTest := struct {
		Arg1 []int
		Arg2 int
		Want int
	}{
		[]int{1, 4, 7, 9, 10},
		9,
		3,
	}
	res := BinarySearch(myTest.Arg1, myTest.Arg2)
	if res != myTest.Want {
		t.Errorf("BinarySearch(%d,%d) == %d, want %d", myTest.Arg1, myTest.Arg2, res, myTest.Want)
	}
}
func TestLeftBound(t *testing.T) {
	myTest := struct {
		Arg1 []int
		Arg2 int
		Want int
	}{
		[]int{1, 4, 4, 4, 7, 9, 10},
		4,
		1,
	}
	res := LeftBound(myTest.Arg1, myTest.Arg2)
	if res != myTest.Want {
		t.Errorf("LeftBound(%d,%d) == %d, want %d", myTest.Arg1, myTest.Arg2, res, myTest.Want)
	}
}
func TestRightBound(t *testing.T) {
	myTest := struct {
		Arg1 []int
		Arg2 int
		Want int
	}{
		[]int{1, 4, 4, 4, 7, 9, 10},
		4,
		3,
	}
	res := RightBound(myTest.Arg1, myTest.Arg2)
	if res != myTest.Want {
		t.Errorf("RightBound(%d,%d) == %d, want %d", myTest.Arg1, myTest.Arg2, res, myTest.Want)
	}
}
