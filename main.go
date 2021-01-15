package main

import (
	"fmt"

	"interview.prep/concepts"
)

func main() {
	// data := []int{1, 2, 3, 2, 5, 4}
	// stringData := "A quick brown fox jumps over a lazy dog"
	// string1 := "a quick"
	// string2 := "quick a"
	// fmt.Printf("Input: %v \n", data)
	// Sorting algorithms - O(n^2)

	// algorithms.InsertionSort(data)
	// algorithms.BubbleSort(data)
	// fmt.Printf("Merge Sort: %v", algorithms.MergeSort(data))
	// algorithms.RandomQuickSort(data, 0, 5)
	// algorithms.QuickSort(data)
	// algorithms.SelectionSort(data)
	node3 := &concepts.ListNode{
		Val: 2,
	}
	node2 := &concepts.ListNode{
		Val:  1,
		Next: node3,
	}
	node1 := &concepts.ListNode{
		Val:  0,
		Next: node2,
	}
	v := concepts.ReverseList(node1)
	for v != nil {
		fmt.Printf("Input: %v, Reverse List (recursion): %v", v.Val)
		v = v.Next
	}
	for v != nil {
		fmt.Printf("Input: %v, Reverse List (recursion): %v", v.Val)
		v = v.Next
	}
}
