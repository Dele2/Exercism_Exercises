package luhn

import (
	"strconv"
    "strings"
)

// function to get characters from strings using index
func getChar(str string, index int) rune {
	return []rune(str)[index]
}

// reverse card string number
func Reverse(card string) string {
	rev := make([]byte, len(card))
	size := len(card)

	for i := range card {
		rev[i] = card[size-1-i]
	}

	return string(rev)
}

// Do the luhn number alcalution
func applyLuhn(idx int, num int) int {

	total := 0
	if isOdd(idx) {
		if num*2 > 9 {
			total += num*2 - 9
		} else {
			total += num * 2
		}
	} else {
		total += num
	}
	return total
}

// check if index is a odd number
func isOdd(num int) bool {
	return num%2 == 1
}

func Valid(id string) bool {

    // Remove all spaces
	card := strings.ReplaceAll(id, " ", "")

    if len(card) < 2 {
		return false
	}
    
	sum := 0

	// Reverse card the numbers
	r_card := Reverse(card)

	for i := range r_card {

		result := string(getChar(r_card, i))

		//convert string to int
		n, err := strconv.Atoi(result)

		//check if error occured
		if err != nil {
			//executes if there is any error
			return false
		} else {
			//executes if there is NO error
			sum += applyLuhn(i, n)
		}
	}
	return sum%10 == 0
}
