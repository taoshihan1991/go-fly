package tools
//划分
func partition(arr *[]int,left int,right int)int{
	privot:=(*arr)[right]
	i:=left-1
	for j:=left;j<right;j++{
		if (*arr)[j]<privot{
			i++
			temp:=(*arr)[i]
			(*arr)[i]=(*arr)[j]
			(*arr)[j]=temp
		}
	}
	temp:=(*arr)[i+1]
	(*arr)[i+1]=(*arr)[right]
	(*arr)[right]=temp
	return i+1
}
//递归
func QuickSort(arr *[]int,left int,right int){
	if left>= right{
		return
	}
	privot:=partition(arr,left,right)
	QuickSort(arr,left,privot-1)
	QuickSort(arr,privot+1,right)
}
//快速排序2
func QuickSort2(arr *[]int,left int,right int){
	if left>= right{
		return
	}
	privot:=(*arr)[left]
	i:=left
	j:=right
	for i<j{
		for i<j && (*arr)[j]>privot{
			j--
		}
		for i<j && (*arr)[i]<=privot{
			i++
		}
		temp:=(*arr)[i]
		(*arr)[i]=(*arr)[j]
		(*arr)[j]=temp
	}
	(*arr)[left]=(*arr)[i]
	(*arr)[i]=privot

	QuickSort(arr,left,i-1)
	QuickSort(arr,i+1,right)
}