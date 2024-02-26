package linkedlist

// https://leetcode.cn/problems/intersection-of-two-linked-lists/description/?envType=study-plan-v2&envId=top-100-liked
//给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。
//如果两个链表不存在相交节点，返回 null

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// GetIntersectionNode 两个链表直接比较，时间复杂度 O(N^2)
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	curr1 := headA
	for curr1 != nil {
		curr2 := headB
		for curr2 != nil {
			if curr1 == curr2 {
				return curr1
			}
			curr2 = curr2.Next
		}
		curr1 = curr1.Next
	}
	return nil
}

// GetIntersectionNode2 记录下来链表每个节点比较
func GetIntersectionNode2(headA, headB *ListNode) *ListNode {
	var l1 []*ListNode
	var l2 []*ListNode
	curr1 := headA
	curr2 := headB
	for curr1 != nil {
		l1 = append(l1, curr1)
		curr1 = curr1.Next
	}
	for curr2 != nil {
		l2 = append(l2, curr2)
		curr2 = curr2.Next
	}
	for i := 0; i < len(l1); i++ {
		for j := 0; j < len(l2); j++ {
			if l1[i] == l2[j] {
				return l1[i]
			}
		}
	}
	return nil
}
