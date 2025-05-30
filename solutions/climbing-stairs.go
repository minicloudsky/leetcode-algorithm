package solutions

/***
假设你正在爬楼梯。需要 n阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

示例 1：

输入：n = 2
输出：2
解释：有两种方法可以爬到楼顶。
1. 1 阶 + 1 阶
2. 2 阶
示例 2：

输入：n = 3
输出：3
解释：有三种方法可以爬到楼顶。
1. 1 阶 + 1 阶 + 1 阶
2. 1 阶 + 2 阶
3. 2 阶 + 1 阶


提示：

1 <= n <= 45
通过次数942,768提交次数1,747,178

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/climbing-stairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// ClimbStairs 先把所有值算出来，直接返回
func ClimbStairs(n int) int {
	climbMap := make(map[int]int)
	climbMap[1] = 1
	climbMap[2] = 2
	for i := 3; i < 50; i++ {
		climbMap[i] = climbMap[i-1] + climbMap[i-2]
	}

	return climbMap[n]
}

// ClimbStairs2 滚动数组
// 0 0 1 1 2
// p q r
func ClimbStairs2(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}
