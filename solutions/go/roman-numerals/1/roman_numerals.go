package romannumerals

import (
    "fmt"
    "strings"
)

func ToRomanNumeral(input int) (string, error) {

    // check bounds
	if input < 1 || input > 3999 {
		return "", fmt.Errorf(" %d is out of range", input)
	}
    
	numList := [...]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	romanList := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}

    count := 0
    var romanNumber string

	for _, v := range numList {
		count = input / v
		romanNumber += strings.Repeat(romanList[v], count)
		input %= v
	}

	return romanNumber, nil
}
