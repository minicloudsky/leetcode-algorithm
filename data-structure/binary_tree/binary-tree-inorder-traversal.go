package binary_tree

// https://leetcode.cn/problems/binary-tree-inorder-traversal/?envType=study-plan-v2&envId=top-100-liked

// 给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。

func InorderTraversal(root *TreeNode) []int {
	var nodes []int
	if root == nil {
		return nodes
	}
	left := InorderTraversal(root.Left)
	right := InorderTraversal(root.Right)
	nodes = append(nodes, left...)
	nodes = append(nodes, root.Val)
	nodes = append(nodes, right...)
	return nodes
}
