// Package weather provide a function to forecast the weather of a city.
package weather

// CurrentCondition Represent the current condition.
var CurrentCondition string

// CurrentLocation Represent the current city.
var CurrentLocation string

// Forecast return the forecast of the city based on the condition it's currently in.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
