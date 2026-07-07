package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	layer_size := len(layers)
	if time < 1 {
		return 2 * layer_size
	}
	return time * layer_size
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodle_weight, sauce_volume := 0, 0.0

	for _, item := range layers {
		if item == "noodles" {
			noodle_weight += 50
		} else if item == "sauce" {
			sauce_volume += 0.2
		}
	}
	return noodle_weight, sauce_volume
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(fLst, myLst []string) {
    // Get the last value from fLst
	value := fLst[len(fLst)-1]
	// Change last value of myLst with value
	myLst[len(myLst)-1] = value
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, num int) []float64 {
	newSlice := make([]float64, len(amounts))
	for i := range amounts {
		amount := (amounts[i] * float64(num) / 2)
		newSlice[i] = amount
	}
	return newSlice
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
