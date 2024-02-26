package single_number

// https://leetcode.cn/problems/single-number/description/

// 在一个整数的无序数组里，除了两个数字以外，其他数字都出现了两次，
// 找到这两个出现一次的数字（百度大搜二面）

// DoubleNumber map记录出现次数，时间复杂度 O(log n)
func DoubleNumber(nums []int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		if _, ok := m[num]; ok {
			m[num]++
		} else {
			m[num] = 1
		}
	}
	var res []int
	for k, v := range m {
		if v == 1 {
			res = append(res, k)
		}
	}

	return res
}

// SingleNumber 两次循环暴力查找
func SingleNumber(nums []int) int {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		found := false
		for j := i + 1; j < length; j++ {
			if nums[i] == nums[j] {
				found = true
			}
			if nums[i] != nums[j] && found == false && j == length-2 {
				return nums[i]
			}
		}
	}
	return -1
}

// SingleNumber2 由于异或运算满足交换律和结合律，所以遍历过程中相同的元素都会异或成0，
// 而成对出现的元素异或结果也为0，因此最终剩下的就是只出现一次的数字
func SingleNumber2(nums []int) int {
	ans := nums[0]
	if len(nums) > 1 {
		for i := 1; i < len(nums); i++ {
			ans = ans ^ nums[i]
		}
	}
	return ans
}
