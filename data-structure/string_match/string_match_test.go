package string_match

import (
	"fmt"
	"testing"
)

func TestStringMatch(t *testing.T) {
	s := "star bucks coffee"
	subStr := "bucks"
	res := StringMatch(s, subStr)
	fmt.Println(res)
}
