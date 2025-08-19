// Package weather provides methods to present weather conditions. 
package weather

// CurrentCondition is stored as string.
var CurrentCondition string
// CurrentLocation is stored as a string.
var CurrentLocation string

// Forecast returns the weather conditions in a city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
