package solutions

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
