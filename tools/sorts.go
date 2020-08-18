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