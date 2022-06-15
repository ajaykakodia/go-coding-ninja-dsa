package main

var StepsTaken int

func bubbleSort(arr []int) {
	l := len(arr)
	swapsHappen := true
	for i := 0; i < l-1; i++ {
		swapsHappen = false
		for j := 0; j < l-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapsHappen = true
			}
			StepsTaken = StepsTaken + 1
		}
		if !swapsHappen {
			break
		}
	}
}
