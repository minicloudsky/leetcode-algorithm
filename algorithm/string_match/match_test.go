package string_match

import (
	"fmt"
	"testing"
)

func TestNaiveStringSearch(t *testing.T) {
	text := "leetcode algorithm"
	pattern := "code"
	pos := NaiveStringSearch(text, pattern)
	fmt.Println(pos)
}
