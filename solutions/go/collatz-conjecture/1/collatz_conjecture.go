package collatzconjecture

import "errors"

func collatz(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return n*3 + 1
}

func CollatzConjecture(n int) (int, error) {
	count := 0
	if n < 1 {
		return count, errors.New("Must be a positive number")
	}

	for n != 1 {
		n = collatz(n)
		count++
	}

	return count, nil
}
