package solutions

import (
	"fmt"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	res := MajorityElement(nums)
	fmt.Println(res)
}

func TestMajorityElement2(t *testing.T) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	res := MajorityElement2(nums)
	fmt.Println(res)
}

func TestMajorityElement3(t *testing.T) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	res := MajorityElement3(nums)
	fmt.Println(res)
}
