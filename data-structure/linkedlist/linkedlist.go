package linkedlist

import "errors"

type ListNode struct {
	Val  int
	Next *ListNode
}

// BuildLinkedList 没有头节点的单链表
func BuildLinkedList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	l := &ListNode{
		Val: values[0],
	}
	curr := l
	for _, val := range values[1:] {
		newNode := &ListNode{
			Val:  val,
			Next: nil,
		}
		curr.Next = newNode
		curr = curr.Next
	}
	return l
}

// InsertLinkedList 头插法，下标从0开始
func InsertLinkedList(head *ListNode, value, position int) (*ListNode, error) {
	newNode := &ListNode{
		Val:  value,
		Next: nil,
	}
	if head == nil || position == 0 {
		newNode.Next = head
		head = newNode
		return head, nil
	}
	current := head
	for i := 1; i < position; i++ {
		if current == nil {
			return nil, errors.New("找不到待插入节点")
		}
		if current.Next == nil && i == position-1 {
			break
		}
		current = current.Next
	}
	newNode.Next = current.Next
	current.Next = newNode
	return head, nil
}

// DeleteLinkedList 下标从0开始
func DeleteLinkedList(head *ListNode, position int) (*ListNode, error) {
	if head == nil {
		return head, nil
	}
	if position == 0 {
		if head.Next == nil {
			head = nil
			return nil, nil
		} else {
			head = head.Next
			return head, nil
		}
	}
	current := head
	pre := current
	for i := 0; i < position; i++ {
		if current == nil {
			return nil, errors.New("找不到待删除节点")
		}
		pre = current
		current = current.Next
	}
	pre.Next = current.Next
	current.Next = nil
	current = nil
	return head, nil
}

func BuildLinkedList2(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	curr := head
	for _, val := range nums[1:] {
		node := &ListNode{Val: val}
		curr.Next = node
		curr = curr.Next
	}

	return head
}
