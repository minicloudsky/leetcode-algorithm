package linkedlist

// 记录下每一个节点，后续遍历到重复节点了，就是有环
// 时间复杂度: O(N) 空间复杂度 O(N) 其中 N 是链表中的节点数
// 有环的会一直遍历下去，遇到重复节点
// 没环的会遍历到尾部，退出
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	m := make(map[*ListNode]*ListNode)
	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = head
		head = head.Next
	}
	return false
}

// 快慢指针
// 快的每次走两步，慢的走一步，走进环里面会相遇
func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast, slow := head.Next, head
	for slow != fast {
		// 当没有环时候退出，有环一定会出现 slow = fast
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

// 龟兔赛跑 快慢指针
// 快的每次走两步，慢的走一步，走进环里面会相遇
// 循环退出条件是 fast 节点或者 fast 下一个节点为空，说明没有环
func hasCycle3(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast, slow := head.Next, head
	for fast != nil && fast.Next != nil {
		if fast == slow {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
