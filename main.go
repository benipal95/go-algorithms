package main

import (
	"fmt"

	"interview.prep/arraysandstrings"
)

func main() {
	data := []int{1, 2, 3, 2, 5, 4}
	// stringData := "A quick brown fox jumps over a lazy dog"
	string1 := "a quick"
	string2 := "quick a"
	fmt.Printf("Input: %v \n", data)
	// Sorting algorithms - O(n^2)

	// algorithms.InsertionSort(data)
	// algorithms.BubbleSort(data)
	// fmt.Printf("Merge Sort: %v", algorithms.MergeSort(data))
	// algorithms.RandomQuickSort(data, 0, 5)
	// algorithms.QuickSort(data)
	// algorithms.SelectionSort(data)

	/* Arrays and Strings coding problems */
	// fmt.Printf("Data: %v \nIsUnique: %v \n", stringData, arraysandstrings.IsUnique(stringData))
	arraysandstrings.CheckPermutation(string1, string2)
}
