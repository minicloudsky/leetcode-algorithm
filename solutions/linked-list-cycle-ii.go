package solutions

// hashtable 记录节点
func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]int)
	for head != nil {
		if _, ok := m[head]; ok {
			return head
		}
		m[head] = 0
		head = head.Next
	}
	return nil
}

/*
方法二：快慢指针，数学题(操场跑圈问题：假设快指针是慢指针速度的2倍，快指针走f步，慢指针走s步，

	f=2s步,并且f=s+nb步(不包含入口节点，b是环型所包含的节点数，n代表跑圈轮数(环形代表圈))，得s=nb，
	而任何到达环形入口节点需走a+nb步(a是第一个节点到环型入口处的所包含的节点数)，故只需在第一次相
	遇时(n=1),快指针移到起始节点和慢指针都走a步即可，此时得到的就是入口节点)
*/
/**
快慢指针的移动：

初始化时，将快指针（fast）和慢指针（slow）都指向链表的头节点（head）。
在每一轮迭代中，慢指针移动一步，快指针移动两步。
如果链表中不存在环，那么快指针会最先到达链表的末尾（fast == nil || fast.Next == nil），此时返回 nil 表示没有环。
如果链表中存在环，那么快指针和慢指针最终会相遇。
找到环的入口：

重新将快指针指向链表的头节点，保持慢指针在相遇点。
两个指针再次以相同的速度移动，每次移动一步。
当两个指针再次相遇时，相遇点即为环的入口。
*/
func detectCycle2(head *ListNode) *ListNode {
	fast := head
	slow := head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			break
		}
	}

	fast = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
