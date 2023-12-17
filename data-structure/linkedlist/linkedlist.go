package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildLinkedList(vals []int) *ListNode {
	l := &ListNode{}
	curr := l
	for _, val := range vals {
		newNode := &ListNode{
			Val:  val,
			Next: nil,
		}
		curr.Next = newNode
		curr = curr.Next
	}
	return l
}
