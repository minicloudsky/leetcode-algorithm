package binary_tree

// https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/?envType=study-plan-v2&envId=top-100-liked

// 给定一个二叉树 root ，返回其最大深度。
// 二叉树的 最大深度 是指从根节点到最远叶子节点的最长路径上的节点数。

// MaxDepth 递归计算左右子树的高度，然后+1
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)
	depth := max(left, right) + 1
	return depth
}
