package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var t, n, m int
	// scan in number of cases, max number of nights, max number gophers
	fmt.Scanf("%d %d %d\n", &t, &n, &m)
	in := bufio.NewReader(os.Stdin)

	// Work out each test case
	for i := 0; i < t; i++ {
		// Test prime integer numbers of blades
		// Then we can use Chinese remainder theorem to find
		// the number?
		for _, div := range []int{2, 3, 5, 7, 11, 13, 17} {
			guess := strings.TrimSpace(strings.Repeat(fmt.Sprintf("%d ", div), 18))
			fmt.Println(guess)
			// Process return
			var total int
			var bladestring string
			bladestring, _ = in.ReadString('\n')
			// blades, _ := in.ReadSlice(' ')
			fmt.Printf("bladestring is %s\n", bladestring)
			blades := strings.Split(bladestring, " ")
			fmt.Println(blades)
			for k := 0; k < 18; k++ {
				num, _ := strconv.Atoi(blades[k])
				total += num
			}
			// find the average
			avg := total / 18
			fmt.Printf("average is %d\n", avg)
		}
	}

}
