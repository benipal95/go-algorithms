package algorithms

import (
	"fmt"
	"math/rand"
	"time"
)

// data = {5, 1, 2, 6, 0, 3, 8, 4}

/*
InsertionSort - O(N^2)
works similar to the way you sort playing cards in your hands.
The array is virtually split into a sorted and an unsorted part.
Values from the unsorted part are picked and placed at the correct position in the sorted part.
*/
func InsertionSort(data []int) {
	for i := 1; i < len(data); i++ {
		j := i - 1
		temp := data[i]
		fmt.Println("temp:", temp)
		for j >= 0 && data[j] > temp {
			fmt.Println("j:", data[j])
			data[j+1] = data[j]
			j--
		}
		fmt.Printf("End of iteration %d and j: %d \n", i, j)
		fmt.Printf("Insertion Sort: %v", data)
		data[j+1] = temp
	}

	// for i := len(data) - 2; i >= 0; i-- {
	// 	temp := data[i]
	// 	j := i + 1
	// 	for j < len(data) && data[j] < temp {
	// 		data[j-1] = data[j]
	// 		j++
	// 	}
	// 	data[j-1] = temp
	// }
	// fmt.Printf("Insertion Sort: %v", data)
}

/*
BubbleSort - O(N^2)
simplest sorting algorithm that works by repeatedly swapping the adjacent elements if they are in wrong order.
*/
func BubbleSort(data []int) {
	for i := 0; i < len(data); i++ {
		for j := 1; j < len(data)-i; j++ {
			if data[j] < data[j-1] {
				data[j-1], data[j] = data[j], data[j-1]
			}
		}
	}

	// Reverse
	// for i := len(data) - 1; i >= 0; i-- {
	// 	for j := len(data) - 2; j >= len(data)-i-1; j-- {
	// 		if data[j] > data[j+1] {
	// 			data[j], data[j+1] = data[j+1], data[j]
	// 		}
	// 	}
	// }
	fmt.Printf("Bubble Sort: %v", data)
}

/*
SelectionSort - O(N^2)
sorts an array by repeatedly finding the minimum element (considering ascending order)
from unsorted part and putting it at the beginning.
The algorithm maintains two subarrays in a given array.
*/
func SelectionSort(data []int) {
	length := len(data)
	for i := 0; i < length; i++ {
		minIndex := 0
		for j := 1; j < length-i; j++ {
			if data[j] > data[minIndex] {
				minIndex = j
			}
		}
		data[minIndex], data[length-i-1] = data[length-i-1], data[minIndex]
	}
	fmt.Printf("Selection Sort: %v", data)
}

/*
MergeSort is a Divide and Conquer algorithm. It divides the input array into two halves,
calls itself for the two halves, and then merges the two sorted halves.
The merge() function is used for merging two halves.
The merge(arr, l, m, r) is a key process that assumes that arr[l..m] and arr[m+1..r] are sorted
and merges the two sorted sub-arrays into one
*/
func MergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	middle := len(data) / 2
	left := MergeSort(data[:middle])
	right := MergeSort(data[middle:])

	return mergeForMergeSort(left, right)

}

func mergeForMergeSort(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	for i := 0; len(left) > 0 || len(right) > 0; i++ {
		if len(left) > 0 && len(right) > 0 {
			if left[0] < right[0] {
				result[i] = left[0]
				left = left[1:]
			} else {
				result[i] = right[0]
				right = right[1:]
			}
		} else if len(left) > 0 {
			result[i] = left[0]
			left = left[1:]
		} else if len(right) > 0 {
			result[i] = right[0]
			right = right[1:]
		}
	}
	return result
}

/*
QuickSort is a Divide and Conquer algorithm. It picks an element as pivot and partitions the given array around the picked pivot.
There are many different versions of quickSort that pick pivot in different ways: Always pick first element as pivot, Always pick last element as pivot,
Pick a random element as pivot, Pick median as pivot. The key process in quickSort is partition(). Target of partitions is, given an array and an element
x of array as pivot, put x at its correct position in sorted array and put all smaller elements (smaller than x) before x, and put all greater elements
(greater than x) after x. All this should be done in linear time.
*/
func QuickSort(data []int) {
	recursionSort(data, 0, len(data)-1)
}

func recursionSort(data []int, left, right int) {
	if left < right {
		pivot := quickSortPartition(data, left, right)
		fmt.Printf("pivot: %v \n", pivot)
		fmt.Printf("calling left side with data: %v \n", data)
		recursionSort(data, left, pivot-1)
		fmt.Printf("calling right side with data: %v \n", data)
		recursionSort(data, pivot+1, right)
	}
}

func quickSortPartition(data []int, left, right int) int {
	for left < right {
		for left < right && data[left] <= data[right] {
			right--
		}
		if left < right {
			data[left], data[right] = data[right], data[left]
			left++
		}
		for left < right && data[left] <= data[right] {
			left++
		}
		if left < right {
			data[left], data[right] = data[right], data[left]
			right--
		}
	}
	return left
}

// RandomQuickSort is the random version
func RandomQuickSort(data []int, start, end int) {
	if end-start > 1 {
		mid := randomPartition(data, start, end)
		RandomQuickSort(data, start, mid)
		RandomQuickSort(data, mid+1, end)
	}
	fmt.Printf("Random Quick Sort: %v", data)
}

func randomPartition(data []int, begin int, end int) int {
	rand.Seed(time.Now().UnixNano())
	i := begin + rand.Intn(end-begin)
	data[i], data[begin] = data[begin], data[i]
	return partition(data, begin, end)
}

func partition(data []int, begin, end int) (i int) {
	cValue := data[begin]
	i = begin
	for j := i + 1; j < end; j++ {
		if data[j] < cValue {
			i++
			data[j], data[i] = data[i], data[j]
		}
	}
	data[i], data[begin] = data[begin], data[i]
	return i
}
