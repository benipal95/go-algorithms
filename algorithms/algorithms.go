package algorithms

import "fmt"

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
}

/*
BubbleSort - O(N^2)
simplest sorting algorithm that works by repeatedly swapping the adjacent elements if they are in wrong order.
*/
func BubbleSort(data []int) {
	for i := 0; i < len(data); i++ {
		for j := 1; j < len(data)-1; j++ {
			if data[j] < data[j-1] {
				data[j-1], data[j] = data[j], data[j-1]
			}
		}
	}
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
		data[length-i-1], data[minIndex] = data[minIndex], data[length-i-1]
	}
	fmt.Printf("Selection Sort: %v", data)
}
