package binary_search

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9, 11}
	target := 9
	res := BinarySearch(nums, target)
	fmt.Println(res)
}

func TestBinarySearch2(t *testing.T) {
	nums := []int{12, 34, 56, 78, 90, 123}
	target := 56
	res := BinarySearch2(nums, target)
	fmt.Println(res)
}
