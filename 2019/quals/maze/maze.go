package main

import (
	"fmt"
)

func main() {
	var t, n int
	var p string

	// scan in number of cases
	fmt.Scanf("%d\n", &t)

	for i := 1; i <= t; i++ {
		// Read in the maze size
		fmt.Scanf("%d\n", &n)

		// Read in Lydia's path
		fmt.Scanf("%s\n", &p)

		// Build answer direct to stdout
		fmt.Printf("Case #%d: ", i)
		for _, d := range p {
			if d == 'S' {
				fmt.Print("E")
			} else {
				fmt.Print("S")
			}
		}
		fmt.Println()
	}
}
