package best_time_to_buy_and_sell_stock

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/?envType=study-plan-v2&envId=top-100-liked
// 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
// 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
// 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

// MaxProfit 暴力，超时。。
func MaxProfit(prices []int) int {
	n := len(prices)
	p := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			profit := prices[j] - prices[i]
			if profit > 0 && profit > p {
				p = profit
			}
		}
	}
	return p
}

// MaxProfit2
// 我们来假设自己来购买股票。随着时间的推移，每天我们都可以选择出售股票与否。那么，假设在第 i 天，如果我们要在今天卖股票，那么我们能赚多少钱呢？
// 显然，如果我们真的在买卖股票，我们肯定会想：如果我是在历史最低点买的股票就好了！太好了，在题目中，我们只要用一个变量记录一个历史
// 最低价格 minprice，我们就可以假设自己的股票是在那天买的。那么我们在第 i 天卖出股票能得到的利润就是 prices[i] - minprice。
// 因此，我们只需要遍历价格数组一遍，记录历史最低点，然后在每一天考虑这么一个问题：如果我是在历史最低点买进的，那么我今天卖出能赚多少钱？
// 当考虑完所有天数之时，我们就得到了最好的答案。
func MaxProfit2(prices []int) int {
	minPrice := 99999999999999999999
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		maxProfit = max(maxProfit, prices[i]-minPrice)
		minPrice = min(minPrice, prices[i])
	}
	return maxProfit
}
