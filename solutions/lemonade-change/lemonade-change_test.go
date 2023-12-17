package lemonade_change

import (
	"fmt"
	"testing"
)

func TestLemonadeChange(t *testing.T) {
	bills := []int{5, 5, 5, 10, 20}
	fmt.Println(lemonadeChange(bills))
	bills2 := []int{5, 5, 10, 10, 20}
	fmt.Println(lemonadeChange(bills2))
	bills3 := []int{5, 5, 10, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 20, 5, 20, 5}
	fmt.Println(lemonadeChange(bills3))
}
