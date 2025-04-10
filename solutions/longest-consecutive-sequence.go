package solutions

// https://leetcode.cn/problems/longest-consecutive-sequence/description/?envType=study-plan-v2&envId=top-100-liked
// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
//
//请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
//示例 1：
//
//输入：nums = [100,4,200,1,3,2]
//输出：4
//解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
//示例 2：
//
//输入：nums = [0,3,7,2,5,8,4,6,0,1]
//输出：9

// LongestConsecutive map记录数字是否存在，存在就一直用map往后判断，记录最大长度
func LongestConsecutive(nums []int) int {
	numMap := make(map[int]bool)
	longestLength := 0
	for _, num := range nums {
		numMap[num] = true
	}

	for num := range numMap {
		if !numMap[num-1] {
			currentNum := num
			currentLength := 1
			for numMap[currentNum+1] {
				currentNum++
				currentLength++
			}
			if currentLength > longestLength {
				longestLength = currentLength
			}
		}
	}
	return longestLength
}
