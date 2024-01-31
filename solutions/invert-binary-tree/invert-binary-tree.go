package invert_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	} else {
		left := root.Left
		root.Left = root.Right
		root.Right = left
		root.Left = invertTree(root.Left)
		root.Right = invertTree(root.Right)
	}
	return root
}
