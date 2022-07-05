package main

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	num := fibonacci(n-1) + fibonacci(n-2)
	return num
}
