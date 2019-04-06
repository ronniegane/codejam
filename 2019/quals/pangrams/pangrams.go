package main

import (
	"fmt"
	"sort"
)

func main() {
	var t, n, l int
	// scan in number of cases
	fmt.Scanf("%d\n", &t)
	var clear string
	var products []int

	for i := 1; i <= t; i++ {
		// Read in N (upper limit of primes) and L (count of products)
		fmt.Scanf("%d %d\n", &n, &l)

		// Read in ciphertext and break into list of products
		products = make([]int, l)
		for j := 0; j < l; j++ {
			fmt.Scanf("%d", &products[j])
		}

		fmt.Printf("products: %v\n", products)

		// Output cleartext
		clear = decode(products)
		fmt.Printf("Case #%d: %s\n", i, clear)
	}
}

func decode(cipher []int) string {
	primes := make([]int, len(cipher)+1)
	// Crack the prime for the second char using the first two products
	primes[1] = gcd(cipher[0], cipher[1])

	primes[0] = cipher[0] / primes[1]

	// Get a list of the single primes
	// primes[x] = cipher[x-1]/primes[x-1]

	for i := 2; i <= len(cipher); i++ {
		primes[i] = cipher[i-1] / primes[i-1]
	}

	fmt.Printf("%v\n", primes)
	// Find uniques and sort
	translator := map[int]string{}
	var uniques []int

	for _, p := range primes {
		// if this prime hasn't been seen before, add it
		if _, ok := translator[p]; !ok {
			translator[p] = ""
			uniques = append(uniques, p)
		}
	}

	fmt.Printf("unique primes: %v\n", uniques)

	// Build up a map of prime to alpha
	sort.Ints(uniques)
	for a, n := range uniques {
		translator[n] = fmt.Sprintf("%c", 65+a)
	}

	fmt.Printf("Dictionary: %v\n", translator)

	// Translate the list of primes
	var clear string
	for _, p := range primes {
		clear += translator[p]
	}
	return clear
}

func gcd(a, b int) int {
	// find the greatest common divisor of a and b
	// uses Euclid's algorithm
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
