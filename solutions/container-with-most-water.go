package solutions

// 盛最多水的容器
// https://leetcode.cn/problems/container-with-most-water/description/?envType=study-plan-v2&envId=top-100-liked
// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 返回容器可以储存的最大水量。
// 说明：你不能倾斜容器。
// 输入：[1,8,6,2,5,4,8,3,7]
// 输出：49
// 解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
// 示例 2：
// 输入：height = [1,1]
// 输出：1
// 提示
// n == height.length
// 2 <= n <= 105
// 0 <= height[i] <= 104
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0
	for left < right {
		width := right - left
		h := 0
		// 在最大容器问题中，移动较低一侧指针的核心原因在于：当前容器的面积受限于较矮的边， 而移动较高的边无法突破这一限制。
		// 通过移动较矮侧的指针，我们试图寻找更高的边，以抵消宽度减少的影响，从而可能获得更大的面积
		// 面积公式与短板效应：面积由 宽度（两指针间距） × 高度（两指针中较矮的值）
		if height[left] < height[right] {
			h = height[left]
			left++
		} else {
			h = height[right]
			right--
		}
		if area := width * h; area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}
