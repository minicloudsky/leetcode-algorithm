package solutions

func isSameTree(p *TreeNode, q *TreeNode) bool {
	var dfs func(node *TreeNode)
	var tree1 []int
	var tree2 []int
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		if r.Left != nil {
			dfs(r.Left)
		}
		tree1 = append(tree1, r.Val)
		if r.Right != nil {
			dfs(r.Right)
		}
	}
	dfs(p)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		if r.Left != nil {
			dfs(r.Left)
		}
		tree2 = append(tree2, r.Val)
		if r.Right != nil {
			dfs(r.Right)
		}
	}
	dfs(q)
	if len(tree1) != len(tree2) {
		return false
	}
	for i := 0; i < len(tree1); i++ {
		if tree1[i] != tree2[i] {
			return false
		}
	}
	return true
}
