package merge_sorted_array

func merge(nums1 []int, m int, nums2 []int, n int) {
	var nums3 []int
	for i := 0; i < len(nums1); i++ {
		nums3 = append(nums3, nums1[i])
	}
	for i := 0; i < len(nums2); i++ {
		nums3 = append(nums3, nums2[i])
	}
	for i := 0; i < len(nums3)-1; i++ {
		for j := i; j < len(nums3); j++ {
			if nums3[i] > nums3[j] {
				tmp := nums3[i]
				nums3[i] = nums3[j]
				nums3[j] = tmp
			}
		}
	}
	nums1 = nums3
}
