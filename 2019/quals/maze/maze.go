package main

import (
	"fmt"
)

func main() {
	var t, n int
	var p, y string

	// scan in number of cases
	fmt.Scanf("%d\n", &t)

	for i := 1; i <= t; i++ {
		// Read in the maze size
		fmt.Scanf("%d\n", &n)

		// Read in Lydia's path
		fmt.Scanf("%s\n", &p)

		// Output answer
		y = maze(n, p)
		fmt.Printf("Case #%d: %s\n", i, y)
	}
}

func maze(n int, p string) string {
	// is the inverse of the string p guaranteed to never cross?
	// paths will only cross if they have the same number of Ss and Es
	// therefore we will be at the same point in the path string
	// and if we use the inverse, we will choose differently to Lydia
	var y string

	for _, d := range p {
		if d == 'S' {
			y += "S"
		} else {
			y += "E"
		}
	}
	return y
}
