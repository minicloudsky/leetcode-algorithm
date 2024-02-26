package binary_search

// BinarySearch 二分查找循环实现
func BinarySearch(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start < end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if target > nums[mid] {
			start = mid
		} else {
			end = mid
		}
	}
	return -1
}

// BinarySearch2 递归实现二分查找
func BinarySearch2(nums []int, target int) int {
	return BinarySearchByRange(nums, target, 0, len(nums)-1)
}

func BinarySearchByRange(nums []int, target, left, right int) int {
	mid := (left + right) / 2
	if target > nums[mid] {
		return BinarySearchByRange(nums, target, mid+1, right)
	} else if target < nums[mid] {
		return BinarySearchByRange(nums, target, left, mid-1)
	} else {
		return mid
	}
}
