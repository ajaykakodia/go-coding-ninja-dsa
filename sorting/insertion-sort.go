package main

func insertionSort(arr []int) {
	l := len(arr)
	for i := 1; i < l; i++ {
		ele := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			StepsTaken = StepsTaken + 1
			if arr[j] > ele {
				arr[j+1] = arr[j]
				continue
			}
			break
		}
		arr[j+1] = ele
	}
}
