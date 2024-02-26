package number_position_in_sorted_array

import (
	"fmt"
	"testing"
)

func TestNumberPositionInSortedArray(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 5, 5, 6, 8, 10, 12}
	target := 5
	num := NumberPositionInSortedArray(nums, target)
	fmt.Println(num)
}

func TestNumberPositionInSortedArray2(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 5, 5, 6, 8, 10, 12}
	target := 5
	num := NumberPositionInSortedArray2(nums, target)
	fmt.Println(num)
}

func TestNumberPositionInSortedArray3(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 5, 5, 6, 8, 10, 12}
	target := 5
	num := NumberPositionInSortedArray3(nums, target)
	fmt.Println(num)
}

func TestNumberPositionInSortedArray4(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 5, 5, 6, 8, 10, 12}
	target := 5
	num := NumberPositionInSortedArray4(nums, target)
	fmt.Println(num)
}
