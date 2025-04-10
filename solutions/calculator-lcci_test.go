package solutions

import (
	"fmt"
	"testing"
)

func TestCalculate(t *testing.T) {
	strings := []string{"1+ 2*3- 4/2"}
	for _, s := range strings {
		fmt.Println(Calculate(s))
	}
}

func TestNum(t *testing.T) {
	var s = "111"
	var num int32
	for _, c := range s {
		if c >= '0' && c <= '9' {
			num = num*10 + c - '0'
		}
	}
	fmt.Println(num)
}
