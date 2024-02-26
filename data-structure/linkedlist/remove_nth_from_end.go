package linkedlist

func getLength(head *ListNode) (length int) {
	for ; head != nil; head = head.Next {
		length++
	}
	return
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	length := getLength(head)
	dummy := &ListNode{Next: head}
	curr := dummy
	for i := 0; i < length-n; i++ {
		curr = curr.Next
	}
	curr.Next = curr.Next.Next
	return dummy.Next
}

// RemoveNthFromEnd2 我们也可以在遍历链表的同时将所有节点依次入栈。根据栈「先进后出」的原则，
// 我们弹出栈的第 n 个节点就是需要删除的节点， 并且目前栈顶的节点就是待删除节点的前驱节点。
// 这样一来，删除操作就变得十分方便了
func RemoveNthFromEnd2(head *ListNode, n int) *ListNode {
	var nodes []*ListNode
	dummy := &ListNode{Next: head}
	curr := dummy
	for ; curr != nil; curr = curr.Next {
		nodes = append(nodes, curr)
	}
	prev := nodes[len(nodes)-n-1]
	prev.Next = prev.Next.Next
	return dummy.Next
}

// RemoveNthFromEnd3 遍历一遍存储每个节点，然后再删掉第 len(nodes) - n 个节点
func RemoveNthFromEnd3(head *ListNode, n int) *ListNode {
	var target *ListNode
	var nodes []*ListNode
	current := head
	for current != nil {
		nodes = append(nodes, current)
		current = current.Next
	}
	// 如果要删除的是头节点，直接返回第二个节点
	if len(nodes) == n {
		return head.Next
	}
	target = nodes[len(nodes)-n]
	current = head
	var prev *ListNode
	for current != nil && current != target {
		prev = current
		current = current.Next
	}

	prev.Next = current.Next
	return head
}

// RemoveNthFromEnd4 双指针，first 先走n步，然后second再走，
// 当first到尾部时，second位置的下一个就是待删除节点
// 时间复杂度: O(L), L 是链表长度
// 空间复杂度: O(1)
func RemoveNthFromEnd4(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	first := head
	second := dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}

func RemoveNthFromEnd5(head *ListNode, n int) *ListNode {
	length := getLength(head)
	prev := head
	for i := 0; i < length-n-1; i++ {
		prev = prev.Next
	}
	// 总长度: length
	// 倒数第n个
	// length - n + 1
	prev.Next = prev.Next.Next
	return head
}
