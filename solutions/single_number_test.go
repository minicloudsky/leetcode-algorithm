package solutions

import (
	"fmt"
	"testing"
)

func TestDoubleNumber(t *testing.T) {
	nums := []int{1, 2, 12, 3, 6, 1, 2, 3, 6, 18}
	res := DoubleNumber(nums)
	fmt.Println(res)
}

func TestSingleNumber(t *testing.T) {
	nums := []int{2, 2, 1}
	res := SingleNumber(nums)
	fmt.Println(res)
}

func TestSingleNumber2(t *testing.T) {
	nums := []int{1, 2, 3, 4, 1, 2, 100, 3, 4}
	res := SingleNumber2(nums)
	fmt.Println(res)
}

func TestOr(t *testing.T) {
	fmt.Println(100 ^ 2 ^ 3 ^ 2 ^ 3 ^ 99 ^ 100)
}
