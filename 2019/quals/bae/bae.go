package main

import (
	"fmt"
	"sort"
	"strings"
)

type SortableList [][]int

func (s SortableList) Len() int {
	return len(s)
}

func (s SortableList) Less(i, j int) bool {
	return s[i][1] < s[j][1]
}

func (s SortableList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	var t, n, b, f int
	// scan in number of test cases
	fmt.Scanf("%d\n", &t)
	// fmt.Printf("test cases: %d\n", t)

	for i := 1; i <= t; i++ {

		// Read in N (total number of workers)
		// B (number of broken workers)
		// and F (max number of queries)
		fmt.Scanf("%d %d %d\n", &n, &b, &f)
		// fmt.Printf("%d %d %d\n", n, b, f)

		known := make([]int, b) // indices of broken workers
		fails := make(SortableList, n)
		// fails is a 2d array: [worker index][failure count]
		for y := 0; y < n; y++ {
			fails[y] = []int{y, 0}
		}
		guess := make([]int, n)
		// first guess is left half 1s, right half 0s
		// then recurse into halves, eg
		// 11110000
		// 11001100
		// 10101010
		// count number of 1s and 0s in response
		// find missing number of 1s and 0s
		// add those numbers to the "fails" count of each index
		// once we have gone F times, find the indices
		// with the largest fail totals
		// these should be the broken workers

		// Loop a maximum of F times
		for loop := 0; loop < f && pow(2, loop+1) <= n; loop++ {

			// Figure out what query to send
			for z := 0; z < n; z++ {
				if z%pow(2, loop+1) < pow(2, loop) {
					guess[z] = 1
				} else {
					guess[z] = 0
				}
			}

			// Send query
			// fmt.Println("query")
			// query := fmt.Sprint(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(guess)), ""), "[]"))
			fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(guess)), ""), "[]"))

			// Process the response
			var resp string
			// wait for response???
			fmt.Scanf("%s\n", &resp)
			// hardcode response to delete index 1000
			// resp = query[:1000] + query[1001:]
			// fmt.Println(resp)

			var ones, zeroes, m1, m0 int
			// counts of 1s and zeroes
			for _, ch := range resp {
				if ch == '1' {
					ones++
				} else {
					zeroes++
				}
			}
			// missing ones and zeroes
			m0 = n/2 - zeroes
			m1 = n/2 - ones
			// fmt.Printf("Missing 0s: %d , 1s: %d\n", m0, m1)

			// Update our failures
			for x, ch := range guess {
				if ch == 1 {
					fails[x][1] += m1
				} else {
					fails[x][1] += m0
				}
			}
			// fmt.Printf("failures: %v\n", fails)

			// build the next query
		}

		// find the workers with the most failures
		// ie the indices of the b largest values
		// in fails
		// maybe make the array 2D, first value original index,
		// second value failscount
		// sort the fails array based on failure count
		// sort.Slice(fails, func(i, j int) bool {
		// 	if fails[i][1] < fails[j][1] {
		// 		return false
		// 	}
		// 	return true
		// })

		sort.Sort(sort.Reverse(fails))

		// fmt.Printf("sorted: %v\n", fails)

		// find the b largest values in array
		for k := 0; k < b; k++ {
			known[k] = fails[k][0]
		}

		// sort it
		sort.Ints(known)
		// send that to the judge as a space-separted string
		fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(known)), " "), "[]"))
	}
}

func pow(base, exp int) int {
	// simple power
	ret := 1
	for i := 0; i < exp; i++ {
		ret *= base
	}
	return ret
}
