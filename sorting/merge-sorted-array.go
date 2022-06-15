package main

func mergeSortedArray(arr1, arr2 []int) []int {
	li := len(arr1)
	lj := len(arr2)
	arr3 := make([]int, 0, li+lj)
	i, j := 0, 0

	for {
		for i < li && arr1[i] < arr2[j] {
			arr3 = append(arr3, arr1[i])
			i++
		}

		if i == li {
			arr3 = append(arr3, arr2[j:]...)
			break
		}
		for j < lj && arr2[j] <= arr1[i] {
			arr3 = append(arr3, arr2[j])
			j++
		}

		if j == lj {
			arr3 = append(arr3, arr1[i:]...)
			break
		}
	}

	return arr3
}
