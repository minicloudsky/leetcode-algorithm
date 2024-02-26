package quick_sort

func QuickSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	// 选择一个基准值
	pivot := nums[0]
	// 分区
	left, right := 0, len(nums)-1
	for i := 1; i <= right; {
		if nums[i] < pivot {
			nums[left], nums[i] = nums[i], nums[left]
			left++
			i++
		} else {
			nums[right], nums[i] = nums[i], nums[right]
			right--
		}
	}
	// 递归地对左右两个分区进行快速排序
	QuickSort(nums[:left])
	QuickSort(nums[left+1:])
}
