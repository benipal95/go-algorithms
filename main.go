package main

import (
	"fmt"

	algo "interview.prep/algorithms"
)

func main() {
	data := []int{5, 1, 2, 6, 0, 3, 8, 4}
	fmt.Printf("Input: %v \n", data)
	// Sorting algorithms - O(n^2)

	// algo.InsertionSort(data)
	// algo.BubbleSort(data)
	algo.SelectionSort(data)
}
