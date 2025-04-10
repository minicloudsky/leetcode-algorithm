package binary_tree

import "testing"

func TestBinaryTree(t *testing.T) {
	array := []int{4, 1, 6, 0, 2, 5, 7, -1, -1, -1, 3, -1, -1, -1, 8}
	root := constructBinaryTree(array)
	printBinaryTree(root, len(array))
}

func TestDfs(t *testing.T) {
	array := []int{4, 1, 6, 0, 2, 5, 7, -1, -1, -1, 3, -1, -1, -1, 8}
	root := constructBinaryTree(array)
	//preOrder(root)
	//inOrder(root)
	postOrder(root)
}
