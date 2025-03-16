package binary_tree

// https://leetcode.cn/problems/symmetric-tree/description/?envType=study-plan-v2&envId=top-100-liked
// 给你一个二叉树的根节点 root ， 检查它是否轴对称。

// 如果一个树的左子树与右子树镜像对称，那么这个树是对称的。
//
//因此，该问题可以转化为：两个树在什么情况下互为镜像？
//
//如果同时满足下面的条件，两个树互为镜像：
//
//它们的两个根结点具有相同的值
//每个树的右子树都与另一个树的左子树镜像对称
//
//
//我们可以实现这样一个递归函数，通过「同步移动」两个指针的方法来遍历这棵树，
//p 指针和 q 指针一开始都指向这棵树的根，随后 p 右移时，q 左移，p 左移时，
//q 右移。每次检查当前 p 和 q 节点的值是否相等，如果相等再判断左右子树是否对称。

func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return check(p.Left, q.Right) && check(p.Right, q.Left) && p.Val == q.Val
}
