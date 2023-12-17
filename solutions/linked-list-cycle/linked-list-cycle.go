package linked_list_cycle

/*
*
https://leetcode.cn/problems/linked-list-cycle/description/?envType=daily-question&envId=2023-12-17

给你一个链表的头节点 head ，判断链表中是否有环。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。注意：pos 不作为参数进行传递 。仅仅是为了标识链表的实际情况。

如果链表中存在环 ，则返回 true 。 否则，返回 false

提示：

链表中节点的数目范围是 [0, 104]
-105 <= Node.val <= 105
pos 为 -1 或者链表中的一个 有效索引 。

进阶：你能用 O(1)（即，常量）内存解决此问题吗？
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针
func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// 数据范围是10000，循环超过1万次，就是有环的
func hasCycle2(head *ListNode) bool {
	count := 0
	for head != nil {
		if count > 10000 {
			return true
		}
		head = head.Next
		count++
	}
	return false
}

// 随便改个数，如果是有环的，下次会再遇到这个数字
func hasCycle3(head *ListNode) bool {
	const val = 34672345875
	for head != nil {
		if head.Val == val {
			return true
		} else {
			head.Val = val
		}
		head = head.Next
	}
	return false
}

// hashtable 记录走过的节点
func hasCycle4(head *ListNode) bool {
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
