package main

func sum_n(n int) int {
	if n == 0 {
		return 0
	}

	return n + sum_n(n-1)
}
