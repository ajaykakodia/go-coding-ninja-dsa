package main

import "fmt"

func main() {
	fmt.Printf("Message from recursion folder")

	n := 7
	factorial := fact(n)
	fmt.Printf("Factorial of %d : %d\n", n, factorial)

	n = 15
	sum := sum_n(n)
	fmt.Printf("Sum of first N natural number of %d : %d\n", n, sum)

	print1ToN(10)
	fmt.Println()

	printNto1(10)
	fmt.Println()

	fib := fibonacci(10)
	fmt.Printf("fibonacci of N  %d : %d\n", 10, fib)

	isSorted := isListSorted([]int{2, 3, 4, 5, 7, 8})
	fmt.Printf("Array is sorted: %v\n", isSorted)

	index := firstIndex([]int{2, 3, 4, 5, 4, 8}, 4)
	fmt.Printf("First index in array: %d\n", index)

	index = lastIndex([]int{2, 3, 4, 5, 4, 8, 10, 13, 12, 4}, 4)
	fmt.Printf("Last index in array: %d\n", index)

	index = lastIndex1([]int{2, 3, 4, 5, 8, 10, 4, 13, 12}, 4)
	fmt.Printf("Last index in array by lastIndex1 : %d\n", index)

	str := replaceCharInString("love bhai ka: ajay", "a", "r")
	fmt.Printf("Last index in array by lastIndex1 : %s\n", str)
}
