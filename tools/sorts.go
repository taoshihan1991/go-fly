package tools

import "sort"

func SortMap(youMap map[string]interface{}) []interface{} {
	keys := make([]string, 0)
	for k, _ := range youMap {
		keys = append(keys, k)
	}
	myMap := make([]interface{}, 0)
	sort.Strings(keys)
	for _, k := range keys {
		myMap = append(myMap, youMap[k])
	}
	return myMap
}

//划分
func partition(arr *[]int, left int, right int) int {
	privot := (*arr)[right]
	i := left - 1
	for j := left; j < right; j++ {
		if (*arr)[j] < privot {
			i++
			temp := (*arr)[i]
			(*arr)[i] = (*arr)[j]
			(*arr)[j] = temp
		}
	}
	temp := (*arr)[i+1]
	(*arr)[i+1] = (*arr)[right]
	(*arr)[right] = temp
	return i + 1
}

//递归
func QuickSort(arr *[]int, left int, right int) {
	if left >= right {
		return
	}
	privot := partition(arr, left, right)
	QuickSort(arr, left, privot-1)
	QuickSort(arr, privot+1, right)
}

//快速排序2
//找到一个基准，左边是所有比它小的，右边是比它大的，分别递归左右
func QuickSort2(arr *[]int, left int, right int) {
	if left >= right {
		return
	}
	privot := (*arr)[left]
	i := left
	j := right
	for i < j {
		for i < j && (*arr)[j] > privot {
			j--
		}
		for i < j && (*arr)[i] <= privot {
			i++
		}
		temp := (*arr)[i]
		(*arr)[i] = (*arr)[j]
		(*arr)[j] = temp
	}
	(*arr)[left] = (*arr)[i]
	(*arr)[i] = privot

	QuickSort(arr, left, i-1)
	QuickSort(arr, i+1, right)
}

//冒泡排序
//比较相邻元素，较大的往右移
func BubbleSort(arr *[]int) {
	flag := true
	lastSwapIndex := 0
	for i := 0; i < len(*arr)-1; i++ {
		sortBorder := len(*arr) - 1 - i
		for j := 0; j < sortBorder; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp := (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
				flag = false
				lastSwapIndex = j
			}
		}
		sortBorder = lastSwapIndex
		if flag {
			break
		}
	}
}

//插入排序
//将未排序部分插入到已排序部分的适当位置
func InsertionSort(arr *[]int) {
	for i := 1; i < len(*arr); i++ {
		curKey := (*arr)[i]
		j := i - 1
		for curKey < (*arr)[j] {
			(*arr)[j+1] = (*arr)[j]
			j--
			if j < 0 {
				break
			}
		}
		(*arr)[j+1] = curKey
	}
}

//选择排序
//选择一个最小值，再寻找比它还小的进行交换
func SelectionSort(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		minIndex := i
		for j := i + 1; j < len(*arr); j++ {
			if (*arr)[j] < (*arr)[minIndex] {
				minIndex = j
			}
		}
		temp := (*arr)[i]
		(*arr)[i] = (*arr)[minIndex]
		(*arr)[minIndex] = temp
	}
}

//归并排序
//合久必分，分久必合，利用临时数组合并两个有序数组
func MergeSort(arr *[]int, left int, right int) {
	if left >= right {
		return
	}

	mid := (left + right) / 2
	MergeSort(arr, left, mid)
	MergeSort(arr, mid+1, right)

	i := left
	j := mid + 1
	p := 0
	temp := make([]int, right-left+1)
	for i <= mid && j <= right {
		if (*arr)[i] <= (*arr)[j] {
			temp[p] = (*arr)[i]
			i++
		} else {
			temp[p] = (*arr)[j]
			j++
		}
		p++
	}

	for i <= mid {
		temp[p] = (*arr)[i]
		i++
		p++
	}
	for j <= right {
		temp[p] = (*arr)[j]
		j++
		p++
	}
	for i = 0; i < len(temp); i++ {
		(*arr)[left+i] = temp[i]
	}
}
