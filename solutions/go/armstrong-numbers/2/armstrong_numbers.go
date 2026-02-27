package armstrong

import "math"



func NumToList (num int) []int {
  
  var numbers []int
  
  for num > 0 {
    numbers = append(numbers, num % 10)
    num /= 10
  }
  return numbers
}

func IsNumber (number int) bool {
  total := 0
  numbers := NumToList(number)
  size := len(numbers)
  
  for _, num := range numbers {
    total += int(math.Pow(float64(num), float64(size)))
  }
  return total == number
}