package algorithms

import "fmt"

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
		fmt.Printf("%v \n", data)
		data[j+1] = temp
	}
}

func BubbleSort(data []int) {
	for i := 0; i < len(data); i++ {
		for j := 1; j < len(data)-1; j++ {
			if data[j] < data[j-1] {
				data[j-1], data[j] = data[j], data[j-1]
			}
		}
	}
	fmt.Printf("%v", data)
}

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
