package main

import "fmt"

// arr = []int{4, 6, 2, 9, 3, 12, 10, 35, 24, 17}Hi Team, Please sone
func quickSort(arr []int) {
	if len(arr) == 1 || len(arr) == 0 {
		return
	}

	pi := pivotIndex(arr)
	fmt.Println(arr)

	quickSort(arr[:pi])
	quickSort(arr[pi+1:])

}

func pivotIndex(arr []int) int {

	ele := arr[0]

	smallNumCount := 0

	for _, val := range arr {
		if ele > val {
			smallNumCount = smallNumCount + 1
		}
	}

	arr[0], arr[smallNumCount] = arr[smallNumCount], arr[0]

	i, j := 0, len(arr)-1

	for i <= j {

		for arr[i] < ele {
			i++
		}

		for arr[j] > ele {
			j--
		}

		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}

	return smallNumCount
}
