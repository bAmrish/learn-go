// Go has `math/rand` packages that provide utility functions to genrate random number.
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
	iteraions := limit * 1000

	for i := 0; i < iteraions; i++ {
		r := rand.Intn(limit)
		_, exists := dist[r]
		if !exists {
			dist[r] = 0
		}
		dist[r] += 1
	}

	fmt.Println(dist)
}
