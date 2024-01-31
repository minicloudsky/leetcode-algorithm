package linkedlist

// ReverseList 遍历一遍节点存下来，重建链表
func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	values := make([]int, 0)
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	newNode := &ListNode{
		Val:  values[len(values)-1],
		Next: nil,
	}
	newHead := newNode
	if len(values) == 1 {
		return newHead
	}
	for i := len(values) - 2; i >= 0; i-- {
		node := &ListNode{
			Val:  values[i],
			Next: nil,
		}
		newNode.Next = node
		newNode = newNode.Next
	}

	return newHead
}
