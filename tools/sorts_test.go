package tools

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr:=[]int{1,8,3,9,4,5,4,7}
	t.Log(arr)
	QuickSort(&arr,0,len(arr)-1)
	t.Log(arr)
}
