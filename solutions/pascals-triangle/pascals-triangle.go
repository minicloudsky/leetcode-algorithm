package pascals_triangle

// https://leetcode.cn/problems/pascals-triangle/description/?envType=study-plan-v2&envId=top-100-liked

// Generate 杨辉三角
func Generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	for i := range ans {
		ans[i] = make([]int, i+1)
		ans[i][i] = 1
		ans[i][0] = 1
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
		}
	}
	return ans
}
