package tools

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{6, 8, 3, 9, 4, 5, 4, 7}
	t.Log(arr)
	QuickSort(&arr, 0, len(arr)-1)
	t.Log(arr)
}
func TestQuickSort2(t *testing.T) {
	arr := []int{6, 8, 3, 9, 4, 5, 4, 7}
	t.Log(arr)
	QuickSort2(&arr, 0, len(arr)-1)
	t.Log(arr)
}
func TestBubbleSort(t *testing.T) {
	arr := []int{6, 8, 3, 9, 4, 5, 4, 7}
	t.Log(arr)
	BubbleSort(&arr)
	t.Log(arr)
}
func TestInsertionSort(t *testing.T) {
	arr := []int{6, 8, 3, 9, 4, 5, 4, 7}
	t.Log(arr)
	InsertionSort(&arr)
	t.Log(arr)
}
func TestSelectionSort(t *testing.T) {
	arr := []int{6, 8, 3, 9, 4, 5, 4, 7}
	t.Log(arr)
	SelectionSort(&arr)
	t.Log(arr)
}
func TestMergeSort(t *testing.T) {
	arr := []int{6, 8, 3, 9, 4, 5, 4, 7}
	t.Log(arr)
	MergeSort(&arr, 0, len(arr)-1)
	t.Log(arr)
}
func TestNilChannel(t *testing.T) {
	NilChannel()
}
