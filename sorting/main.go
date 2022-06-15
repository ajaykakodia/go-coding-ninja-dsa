package main

import "fmt"

func main() {
	fmt.Println("Executing main from sorting folder")
	fmt.Println("------------ Bubble Sort --------------")
	arr := []int{4, 6, 2, 9, 3}
	bubbleSort(arr)
	fmt.Println(arr)
	fmt.Println("Steps Taken in Bubble Sort : ", StepsTaken)
	StepsTaken = 0
	arr = []int{1, 2, 3, 4, 5}
	bubbleSort(arr)
	fmt.Println(arr)
	fmt.Println("Steps Taken in Bubble Sort : ", StepsTaken)

	fmt.Println("------------ Insertion Sort --------------")
	StepsTaken = 0
	arr = []int{4, 6, 2, 9, 3}
	insertionSort(arr)
	fmt.Println(arr)
	fmt.Println("Steps Taken in Insertion Sort : ", StepsTaken)
	StepsTaken = 0
	arr = []int{1, 2, 3, 4, 5}
	insertionSort(arr)
	fmt.Println(arr)
	fmt.Println("Steps Taken in Insertion Sort : ", StepsTaken)

	fmt.Println("------------ Selection Sort --------------")
	StepsTaken = 0
	arr = []int{4, 6, 2, 9, 3}
	selectionSort(arr)
	fmt.Println(arr)
	fmt.Println("Steps Taken in Selection Sort : ", StepsTaken)
	StepsTaken = 0
	arr = []int{1, 2, 3, 4, 5}
	selectionSort(arr)
	fmt.Println(arr)
	fmt.Println("Steps Taken in Selection Sort : ", StepsTaken)

	fmt.Println("------------ Merge Sorted Array --------------")
	arr1 := []int{4, 7, 10, 15, 20, 38, 39}
	arr2 := []int{5, 8, 11, 14, 22, 35, 37}
	mergeSortedArray(arr1, arr2)
	fmt.Println(mergeSortedArray(arr1, arr2))

	fmt.Println("------------ Merge Sort --------------")
	arr = []int{4, 6, 2, 9, 3, 12, 10, 35, 24, 17}
	arr = mergeSort(arr)
	fmt.Println(arr)

	fmt.Println("------------ Quick Sort --------------")
	arr = []int{4, 6, 2, 9, 3, 12, 10, 35, 24, 17}
	quickSort(arr)
	fmt.Println(arr)
}
