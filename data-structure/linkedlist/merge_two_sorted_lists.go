package linkedlist

// MergeTwoLists 递归 将两个链表里面值小的和另一个链表合并
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val < list2.Val {
		list1.Next = MergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = MergeTwoLists(list1, list2.Next)
		return list2
	}
}

// MergeTwoLists2 新建一个链表，将两个链表合并过来
func MergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	newList := &ListNode{}
	curr := newList
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			newNode := &ListNode{Val: list1.Val}
			curr.Next = newNode
			curr = curr.Next
			list1 = list1.Next
		} else {
			newNode := &ListNode{Val: list2.Val}
			curr.Next = newNode
			curr = curr.Next
			list2 = list2.Next
		}
	}
	// 遍历完只剩下一个链表不为空了
	if list1 != nil {
		curr.Next = list1
	} else {
		curr.Next = list2
	}
	return newList.Next
}
