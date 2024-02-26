package number_position_in_sorted_array

import "sort"

// 给定一个数字，返回它在有序数组中出现的次数(百度大搜二面)

// NumberPositionInSortedArray  遍历整个有序数组，O(N)
func NumberPositionInSortedArray(nums []int, target int) int {
	count := 0
	for _, num := range nums {
		if num == target {
			count++
		}
	}
	return count
}

// NumberPositionInSortedArray2  二分查找，找到后左右移动，找到所有等于target的值计算次数
// 时间复杂度: O(log n + k), k是等于target的数字个数
func NumberPositionInSortedArray2(nums []int, target int) int {
	length := len(nums)
	pos := BinarySearch(nums, 0, length-1, target)
	num := 0
	for i := pos; nums[i] == target && i <= length-1; i++ {
		num++
	}
	for i := pos - 1; nums[i] == target && i >= 0; i-- {
		num++
	}
	return num
}

func BinarySearch(nums []int, left, right, target int) int {
	mid := (left + right) / 2
	if nums[mid] < target {
		return BinarySearch(nums, 0, mid-1, target)
	} else if nums[mid] > target {
		return BinarySearch(nums, mid+1, right, target)
	}
	return mid
}

// NumberPositionInSortedArray3 二分查找，找到左右边界 时间复杂度: O(log n)
func NumberPositionInSortedArray3(nums []int, target int) int {
	leftIdx := binarySearchFindPosition(nums, target, true)
	if leftIdx == len(nums) || nums[leftIdx] != target {
		return 0
	}
	rightIdx := binarySearchFindPosition(nums, target, false)
	return rightIdx - leftIdx
}

func binarySearchFindPosition(nums []int, target int, left bool) int {
	low, high := 0, len(nums)
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] > target || (left && nums[mid] == target) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

// NumberPositionInSortedArray4 标准库
func NumberPositionInSortedArray4(nums []int, target int) int {
	leftIdx := sort.SearchInts(nums, target)
	if leftIdx == len(nums) || nums[leftIdx] != target {
		return 0
	}
	rightIdx := sort.SearchInts(nums, target+1)
	return rightIdx - leftIdx
}
