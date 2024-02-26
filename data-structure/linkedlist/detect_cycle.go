package linkedlist

// https://leetcode.cn/problems/linked-list-cycle-ii/description/?envType=study-plan-v2&envId=top-100-liked
// 给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

// DetectCycle map存储节点，有环会再次遇到
func DetectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := m[head]; ok {
			return head
		}
		m[head] = struct{}{}
		head = head.Next
	}
	return nil
}

// DetectCycle2 快慢指针 我们使用两个指针，fast 与 slow。它们起始都位于链表的头部。随后，slow 指针每次向后移动一个位置，
// 而 fast 指针向后移动两个位置。如果链表中存在环，则 fast 指针最终将再次与 slow 指针在环中相遇。
// 当发现 slow 与 fast 相遇时，我们再额外使用一个指针 ptr 。起始，它指向链表头部；随后，它和 slow 每次向后移动一个位置。最终，它们会在入环点相遇
func DetectCycle2(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if slow == fast {
			ptr := head
			for ptr != slow {
				ptr = ptr.Next
				slow = slow.Next
			}
			return ptr
		}
	}
	return nil
}
