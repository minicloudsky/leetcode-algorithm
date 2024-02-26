package best_time_to_buy_and_sell_stock

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	//prices := []int{7, 1, 5, 3, 6, 4}
	prices := []int{7, 6, 4, 3, 1}
	profit := MaxProfit(prices)
	fmt.Println(profit)
}
