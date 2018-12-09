package main

import (
	"fmt"
)

func main() {
	// read in number of test cases
	var t int
	fmt.Scanf("%d\n", &t)

	for i := 0; i < t; i++ {
		var a, b, guess, maxTries int
		// Read in limits
		fmt.Scanf("%d %d\n", &a, &b)
		// lower limit is exclusive
		a++
		// Read in max number of tries
		fmt.Scanf("%d\n", &maxTries)
		for command := ""; command != "CORRECT"; fmt.Scanf("%s\n", &command) {
			switch command {
			case "TOO_BIG":
				// Replace upper limit with guess
				b = guess - 1
			case "TOO_SMALL":
				a = guess + 1
			case "WRONG_ANSWER":
				return
			} // return the midpoint of the two numbers as the new guess
			guess = (a + b) / 2
			fmt.Printf("%d\n", guess)
		}
	}
}
