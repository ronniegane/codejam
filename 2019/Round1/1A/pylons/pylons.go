package main

import "fmt"

func main() {
	var t, r, c int
	// scan in number of cases
	fmt.Scanf("%d\n", &t)

	// Work out each test case
	for i := 0; i < t; i++ {
		fmt.Scanf("%d %d\n", &r, &c)
		fmt.Printf("Case #%d: %s\n", i+1, solve(r, c))
	}

}

func solve(r, c int) string {
	// Returns either IMPOSSIBLE
	// Or POSSIBLE, followed by the moves to make (r*c total lines)
	// in format "row col"
	// rows and cols are 1-indexed

	// Impossible if grid is 2x2, 2x3 or 3x3 as there is at least
	// one cell that is in view of _all_ others
	if r <= 3 || c <= 3 {
		return "IMPOSSIBLE"
	}

	// Is a depth-first-search going to be too inefficient?
	// Have to choose a starting cell
	// Build up a list of visited cells
	// for each step, choose one from the remaining cells
	// that is not on a diagonal or same row or column

	return "IMPOSSIBLE"
}
