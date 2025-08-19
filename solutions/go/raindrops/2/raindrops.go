package raindrops

import "strconv"

func Convert(number int) string {

	result := ""
	nums := []int{3, 5, 7}

	messages := map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}

	for _, num := range nums {
		if number%num == 0 {
			result += messages[num]
		}
	}

	if len(result) == 0 {
		return strconv.Itoa(number)
	}

	return result
}