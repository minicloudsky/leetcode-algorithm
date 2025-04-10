package solutions

import (
	"fmt"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	array := []int{4, 1, 6, 0, 2, 5, 7, -1, -1, -1, 3, -1, -1, -1, 8}
	root := constructBinaryTree(array)
	res := levelOrder(root)
	fmt.Println(res)
}

func TestLevelOrder2(t *testing.T) {
	array := []int{4, 1, 6, 0, 2, 5, 7, -1, -1, -1, 3, -1, -1, -1, 8}
	root := constructBinaryTree(array)
	var currentNodes []*TreeNode
	var ans [][]int
	currentNodes = append(currentNodes, root)
	for currentNodes != nil {
		var nextNodes []*TreeNode
		var vals []int
		for _, node := range currentNodes {
			vals = append(vals, node.Val)
			if node.Left != nil {
				nextNodes = append(nextNodes, node.Left)
			}
			if node.Right != nil {
				nextNodes = append(nextNodes, node.Right)
			}
		}
		ans = append(ans, vals)
		currentNodes = nextNodes
	}
	fmt.Println(ans)
}
