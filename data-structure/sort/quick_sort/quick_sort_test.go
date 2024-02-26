package quick_sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := []int{12, 8, 2, 1, 89, 76, 31, 47}
	QuickSort(nums)
	fmt.Println(nums)
}
