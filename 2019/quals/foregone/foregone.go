package main

import "fmt"

/*
At least one of the digits of N is a 4.
*/

func main() {
	var t, n, a, b int
	// scan in number of cases
	fmt.Scanf("%d\n", &t)

	for i := 1; i <= t; i++ {
		// Read in target number
		fmt.Scanf("%d\n", &n)

		// Output potential numbers
		a, b = nofours(n)
		fmt.Printf("Case #%d: %d %d\n", i, a, b)
	}
}

func nofours(n int) (int, int) {
	// return a and b such that a + b = n
	// and no
	// easy option is to return a where 4 replaced with 3
	// and b with 1 in the position of the 4s in n
	var a, b, dig, t int
	t = 1 // power of 10
	for n > 0 {
		dig = n % 10
		if dig == 4 {
			a += t * 3
			b += t * 1
		} else {
			a += t * dig
		}
		n = n / 10
		t *= 10
	}
	return a, b
}
