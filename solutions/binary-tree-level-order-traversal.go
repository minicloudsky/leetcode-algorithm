package solutions

// https://leetcode.cn/problems/binary-tree-level-order-traversal/description/
// 给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[3],[9,20],[15,7]]
// 示例 2：
//
// 输入：root = [1]
// 输出：[[1]]
// 示例 3：
//
// 输入：root = []
// 输出：[]
//
// 提示：
//
// 树中节点数目在范围 [0, 2000] 内
// -1000 <= Node.val <= 1000
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ans [][]int
	var currentNodes []*TreeNode
	currentNodes = append(currentNodes, root) // 根节点开始遍历每一层
	for currentNodes != nil {
		var currAns []int
		var nextNodes []*TreeNode
		for _, node := range currentNodes {
			currAns = append(currAns, node.Val) //把这一层的数据加进去
			if node.Left != nil {
				nextNodes = append(nextNodes, node.Left) //把下一层要遍历的节点加进去
			}
			if node.Right != nil {
				nextNodes = append(nextNodes, node.Right)
			}
		}
		ans = append(ans, currAns)
		currentNodes = nextNodes
	}
	return ans
}
