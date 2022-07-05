package main

import "fmt"

// print1ToN - Print 1 2 3 4 5 6 ... N
func print1ToN(n int) {
	if n == 0 {
		return
	}
	print1ToN(n - 1)
	fmt.Printf("%-3d", n)
}

// printNto1 - Print N ... 5 4 3 2 1
func printNto1(n int) {
	if n == 0 {
		return
	}
	fmt.Printf("%-3d", n)
	printNto1(n - 1)
}
