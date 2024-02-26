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

func TestInsertLinkedList(t *testing.T) {
	vals := []int{1, 2, 3, 4, 5}
	list := BuildLinkedList(vals)
	position := 0
	value := 888
	list, err := InsertLinkedList(list, value, position)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(list)
}

func TestDeleteLinkedList(t *testing.T) {
	vals := []int{1, 2, 3, 4, 5}
	list := BuildLinkedList(vals)
	position := 2
	list, err := DeleteLinkedList(list, position)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(list)
}
