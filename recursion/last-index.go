package main

func lastIndex(arr []int, ele int) int {
	if len(arr) == 0 {
		return -1
	}
	if arr[len(arr)-1] == ele {
		return len(arr) - 1
	}
	return lastIndex(arr[:len(arr)-1], ele)
}

func lastIndex1(arr []int, ele int) int {
	if len(arr) == 0 {
		return -1
	}
	index := lastIndex(arr[1:], ele)
	if index == -1 {
		if arr[0] == ele {
			return 0
		}
		return index
	}
	return 1 + index
}
