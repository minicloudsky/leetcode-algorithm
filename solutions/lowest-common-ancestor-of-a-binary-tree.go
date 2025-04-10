package solutions

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/description/
// 二叉树的最近公共祖先
// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
//
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

// dfs遍历，记录父节点，然后找到p和q 公共的父节点
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	parent := make(map[int]*TreeNode)
	visited := make(map[int]bool)
	var dfs = func(r *TreeNode) {}
	dfs = func(r *TreeNode) { // 标记下父节点
		if r == nil {
			return
		}
		if r.Left != nil {
			parent[r.Left.Val] = r
			dfs(r.Left)
		}
		if r.Right != nil {
			parent[r.Right.Val] = r
			dfs(r.Right)
		}
	}
	dfs(root)
	for p != nil { // 标记下p访问过的节点
		visited[p.Val] = true
		p = parent[p.Val]
	}
	for q != nil {
		if visited[q.Val] { // p 和 q 都共有的父节点
			return q
		}
		q = parent[q.Val]
	}
	return nil
}
