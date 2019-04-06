package main

import "fmt"

func main() {
	var t int
	// scan in number of cases
	fmt.Scanf("%d\n", &t)

	for i := 1; i <= t; i++ {
		// Read in target variables from line
		fmt.Scanf("%d\n", &n)

		// Output answer
		a, b = myfunc(n)
		fmt.Printf("Case #%d: %d %d\n", i, a, b)
	}
}
