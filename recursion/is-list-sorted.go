package main

func isListSorted(arr []int) bool {
	if len(arr) == 1 {
		return true
	}
	if arr[0] > arr[1] {
		return false
	}
	return isListSorted(arr[1:])
}
