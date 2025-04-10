package solutions

import (
	"fmt"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	arr := []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}
	root := constructBinaryTree(arr)
	fmt.Println(root)
}
