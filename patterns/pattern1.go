package main

import "fmt"

/*
	****
	****
	****
	****
	Print pattern for any given number n=4 as above
*/
func Pattern1(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
	1111
	2222
	3333
	4444
	Print pattern for any given number n=4 as above
*/
func Pattern2(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(i + 1)
		}
		fmt.Println()
	}
}

/*
	1234
	1234
	1234
	1234
	Print pattern for any given number n=4 as above
*/
func Pattern3(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(j + 1)
		}
		fmt.Println()
	}
}

/*
	4321
	4321
	4321
	4321
	Print pattern for any given number n=4 as above
*/
func Pattern4(n int) {
	for i := 0; i < n; i++ {
		for j := n; j > 0; j-- {
			fmt.Print(j)
		}
		fmt.Println()
	}
}

/*
	1
	12
	123
	1234
	Print pattern for any given number n=4 as above
*/
func Pattern5(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(j + 1)
		}
		fmt.Println()
	}
}

/*
	1
	23
	345
	4567
	Print pattern for any given number n=4 as above
*/
func Pattern6(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(i + j + 1)
		}
		fmt.Println()
	}
}

/*
	1
	23
	456
	78910
	Print pattern for any given number n=4 as above
*/
func Pattern7(n int) {
	e := 0
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			e = e + 1
			fmt.Print(e)

		}
		fmt.Println()
	}
}

/*
	ABCD
	ABCD
	ABCD
	ABCD
	Print pattern for any given number n=4 as above
*/
func Pattern8(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(string('A' + j))
		}
		fmt.Println()
	}
}

/*
	ABCD
	BCDE
	CDEF
	DEFG
	Print pattern for any given number n=4 as above
*/
func Pattern9(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(string('A' + i + j))
		}
		fmt.Println()
	}
}

/*
	****
	***
	**
	*
	Print pattern for any given number n=4 as above
*/
func Pattern10(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
	   *
	  **
	 ***
	****
	Print pattern for any given number n=4 as above
*/
func Pattern11(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j >= (n - i - 1) {
				fmt.Print("*")
				continue
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

/*
	   1
	  12
	 123
	1234
	Print pattern for any given number n=4 as above
*/
func Pattern12(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
	   1
	  121
	 12321
	1234321
	Print pattern for any given number n=4 as above
*/
func Pattern13(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print(j + 1)
		}
		for j := 0; j <= i-1; j++ {
			fmt.Print(i - j)
		}
		fmt.Println()
	}
}
