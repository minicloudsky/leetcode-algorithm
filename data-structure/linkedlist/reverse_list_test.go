package linkedlist

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	list := BuildLinkedList(values)
	newList := ReverseList(list)
	fmt.Println(newList)
}

func TestReverseList2(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	list := BuildLinkedList(values)
	newList := ReverseList2(list)
	fmt.Println(newList)
}

func TestReverseList3(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	list := BuildLinkedList(values)
	newList := ReverseList3(list)
	fmt.Println(newList)
}

func TestReverseList4(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	list := BuildLinkedList(values)
	newList := ReverseList4(list)
	fmt.Println(newList)
}

func TestReverseList5(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	list := BuildLinkedList2(values)
	newList := ReverseList5(list)
	fmt.Println(newList)
}
