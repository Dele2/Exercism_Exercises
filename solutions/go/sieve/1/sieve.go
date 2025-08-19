package sieve

import (
	"math"
)

func Sieve(limit int) []int {
    
	primeBooleans := make([]bool, limit+1)

	// everyone starts as false by default, now set everything from 2 onwards equal to true.
	for p := 2; p <= limit; p++ {
		primeBooleans[p] = true
	}

	for p := 2; float64(p) <= math.Sqrt(float64(limit)); p++ {

		// is 'p' prime? If so, cross off it's multiples.
		if primeBooleans[p] {
			primeBooleans = CrossOfMultiples(primeBooleans, p)
		}
	}

    // first, create a slice of length 0
	primeList := make([]int, 0)
    
    for p := range primeBooleans {
		if primeBooleans[p] {
			// append 'p' to our list
			primeList = append(primeList, p)
		}
	}

	return primeList
}

func CrossOfMultiples(primeBooleans []bool, p int) []bool {

	n := len(primeBooleans) - 1

	// consider every multiple of 'p', starting with '2p', and "cross it off" by setting it's
	// corresponding entry of the slice to false
	for k := 2 * p; k <= n; k += p {
		primeBooleans[k] = false
	}

	return primeBooleans
}
