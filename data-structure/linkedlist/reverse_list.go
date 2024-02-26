package linkedlist

// ReverseList 迭代法: 记录前一个节点和后一个节点，遍历时候改为向前的指针，时间复杂度O(N)
func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// ReverseList2 递归法
func ReverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := ReverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// ReverseList3 头插法插入新链表
func ReverseList3(head *ListNode) *ListNode {
	var newHead *ListNode
	for head != nil {
		newNode := &ListNode{Val: head.Val}
		newNode.Next = newHead
		newHead = newNode
		head = head.Next
	}
	return newHead
}

// ReverseList4 遍历一遍节点存下来，重建链表
func ReverseList4(head *ListNode) *ListNode {
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

func ReverseList5(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
