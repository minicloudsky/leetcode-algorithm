package linkedlist

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	values1 := []int{1, 2, 3, 4, 5}
	values2 := []int{6, 7, 8, 9, 10}
	list1 := BuildLinkedList(values1)
	list2 := BuildLinkedList(values2)
	newList := MergeTwoLists(list1, list2)
	fmt.Println(newList)
}

func TestMergeTwoLists2(t *testing.T) {
	values1 := []int{1, 2, 3, 4, 5}
	values2 := []int{6, 7, 8, 9, 10}
	list1 := BuildLinkedList(values1)
	list2 := BuildLinkedList(values2)
	newList := MergeTwoLists2(list1, list2)
	fmt.Println(newList)
}
