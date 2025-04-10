package solutions

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func TestTwoSum2(t *testing.T) {
	nums := []int{2, 7, 11, 15, 13, 21}
	target := 32
	fmt.Println(twoSum2(nums, target))
}
