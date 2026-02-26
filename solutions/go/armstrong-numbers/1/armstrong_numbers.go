package armstrong

import (
	"strconv"
	"math"
)

func IsNumber(n int) bool {
	str := strconv.Itoa(n)
	size := len(str)
	total := 0

	runes := []rune(str) // Convert string to a slice of runes

    for i := 0; i < len(runes); i++ {
        total += int(math.Pow(float64(runes[i] - '0'), float64(size)))
    }

	return total == n
}
