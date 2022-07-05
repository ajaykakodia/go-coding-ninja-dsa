package main

func firstIndex(arr []int, ele int) int {
	if len(arr) == 0 {
		return -1
	}
	if ele == arr[0] {
		return 0
	}
	index := firstIndex(arr[1:], ele)
	if index == -1 {
		return index
	}

	return 1 + index
}
