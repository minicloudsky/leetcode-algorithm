package majority_element

import "sort"

// https://leetcode.cn/problems/majority-element/description/?envType=study-plan-v2&envId=top-100-liked
// 给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
//
//你可以假设数组是非空的，并且给定的数组总是存在多数元素。
//
//示例 1：
//
//输入：nums = [3,2,3]
//输出：3
//示例 2：
//
//输入：nums = [2,2,1,1,1,2,2]
//输出：2
//
//提示：
//n == nums.length
//1 <= n <= 5 * 104
//-109 <= nums[i] <= 109
//
//进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。

// MajorityElement hashtable 记录出现次数
func MajorityElement(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		if _, ok := m[num]; ok {
			m[num]++
		} else {
			m[num] = 1
		}
	}
	for k, v := range m {
		if v > len(nums)/2 {
			return k
		}
	}
	return -1
}

// MajorityElement2 如果将数组 nums 中的所有元素按照单调递增或单调递减的顺序排序，
// 那么下标为 2/n 的元素一定是众数
func MajorityElement2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// MajorityElement3 Boyer-Moore 投票算法
// 如果我们把众数记为 +1+1+1，把其他数记为 −1-1−1，将它们全部加起来，
// 显然和大于 0，从结果本身我们可以看出众数比其他数多。
func MajorityElement3(nums []int) int {
	count := 0
	candidate := 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count += 1
		} else {
			count -= 1
		}
	}
	return candidate
}
