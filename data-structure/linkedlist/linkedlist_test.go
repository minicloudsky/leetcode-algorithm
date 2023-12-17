package linkedlist

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	vals := []int{1, 2, 3, 4, 5}
	list := BuildLinkedList(vals)
	fmt.Println(list)
}
