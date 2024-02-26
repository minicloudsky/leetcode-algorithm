package binary_tree

// https://leetcode.cn/problems/invert-binary-tree/?envType=study-plan-v2&envId=top-100-liked

// 给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。

// InvertTree 递归反转左右子树
func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	left := root.Left
	root.Left = root.Right
	root.Right = left
	root.Left = InvertTree(root.Left)
	root.Right = InvertTree(root.Right)

	return root
}
