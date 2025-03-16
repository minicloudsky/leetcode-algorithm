package longest_consecutive_sequence

import (
	"fmt"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	nums := []int{100, 4, 200, 1, 3, 2}
	length := LongestConsecutive(nums)
	fmt.Println(length)
}
