package main

func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	mid := len(arr) / 2
	arr1 := mergeSort(arr[0:mid])
	arr2 := mergeSort(arr[mid:])
	return mergeSortedArray(arr1, arr2)
}
