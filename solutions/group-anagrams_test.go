package solutions

import (
	"fmt"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := GroupAnagrams(strs)
	fmt.Println(res)
}
