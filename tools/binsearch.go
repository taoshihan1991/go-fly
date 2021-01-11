package tools

func BinarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1 //注意

	for left <= right { //注意
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1 //注意
		} else if nums[mid] > target {
			right = mid - 1 //注意
		}
	}
	return -1
}
func LeftBound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left := 0
	right := len(nums) //注意

	for left < right { //注意
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid //注意
		}
	}
	if left == len(nums) || nums[left] != target {
		return -1
	}
	return left
}
func LeftBound2(nums []int, target int) int {
	left := 0
	right := len(nums) - 1 //注意

	for left <= right { //注意
		mid := left + (right-left)/2
		if nums[mid] == target {
			//收缩右侧边界
			right = mid - 1
		} else if nums[mid] < target {
			//搜索区间变为 [mid+1, right]
			left = mid + 1 //注意
		} else if nums[mid] > target {
			//搜索区间变为 [left, mid-1]
			right = mid - 1
		}
	}
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}
func RightBound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1 //注意

	for left <= right { //注意
		mid := left + (right-left)/2
		if nums[mid] == target {
			//收缩左侧边界
			left = mid + 1
		} else if nums[mid] < target {
			//搜索区间变为 [mid+1, right]
			left = mid + 1 //注意
		} else if nums[mid] > target {
			//搜索区间变为 [left, mid-1]
			right = mid - 1
		}
	}
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}
