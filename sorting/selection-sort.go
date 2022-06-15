package main

func selectionSort(arr []int) {
	l := len(arr)

	for i := 0; i < l-1; i++ {
		svi := i
		for j := i + 1; j < l; j++ {
			if arr[svi] > arr[j] {
				svi = j
			}
			StepsTaken++
		}
		arr[i], arr[svi] = arr[svi], arr[i]
	}
}
