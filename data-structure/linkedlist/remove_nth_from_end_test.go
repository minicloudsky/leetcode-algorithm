package linkedlist

import (
	"fmt"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	values := []int{0, 1, 2, 3, 4, 5}
	n := 4
	list := BuildLinkedList(values)
	newList := RemoveNthFromEnd(list, n)
	fmt.Println(newList)
}

func TestRemoveNthFromEnd2(t *testing.T) {
	values := []int{0, 1, 2, 3, 4, 5}
	n := 4
	list := BuildLinkedList(values)
	newList := RemoveNthFromEnd2(list, n)
	fmt.Println(newList)
}

func TestRemoveNthFromEnd3(t *testing.T) {
	values := []int{0, 1, 2, 3, 4, 5}
	n := 4
	list := BuildLinkedList(values)
	newList := RemoveNthFromEnd3(list, n)
	fmt.Println(newList)
}

func TestRemoveNthFromEnd4(t *testing.T) {
	values := []int{0, 1, 2, 3, 4, 5}
	n := 6
	list := BuildLinkedList(values)
	newList := RemoveNthFromEnd4(list, n)
	fmt.Println(newList)
}

func TestRemoveNthFromEnd5(t *testing.T) {
	values := []int{0, 1, 2, 3, 4, 5}
	n := 6
	list := BuildLinkedList(values)
	newList := RemoveNthFromEnd5(list, n)
	fmt.Println(newList)
}
