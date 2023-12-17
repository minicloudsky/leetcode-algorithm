package delete_greatest_value_in_each_row

import (
	"fmt"
	"testing"
)

func TestDeleteGreatestValue(t *testing.T) {
	grid := [][]int{{1, 2, 4}, {3, 3, 1}}
	fmt.Println(deleteGreatestValue(grid))
	grid2 := [][]int{{10}}
	fmt.Println(deleteGreatestValue(grid2))
}
