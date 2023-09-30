// Go has `math/rand` packages that provide utility functions to generate random number.
// This program is an attempt to capture the distribution of those random numbers over different ranges.
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Random distribution")
	distribution(3)
	distribution(10)
	distribution(100)
}

func distribution(limit int) {
	dist := map[int]int{}
	iterations := limit * 1000

	for i := 0; i < iterations; i++ {
		r := rand.Intn(limit)
		// _, exists := dist[r]
		// if !exists {
		// 	dist[r] = 0
		// }
		// If a key does not exists in a map the zero-value of
		// value type is returned. In this case the value type is int
		// with the zero-value of 0, so when we set the key for the first
		// time we are just setting it to 1.
		// So the above check is not required
		dist[r] += 1
	}

	fmt.Println(dist)
}
