package bubblesort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{10, 23, 45, 1, 7, 12, -1, -5, 34, 7}
	arr2 := BubbleSort(arr)
	fmt.Println(arr2)
}
