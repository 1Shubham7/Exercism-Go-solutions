// Package weather provides info about the weather forecast.
package weather

// CurrentCondition represents the current condition.
var CurrentCondition string
// CurrentLocation represents the current Location.
var CurrentLocation string

// Forecast function returns a string for the weather forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
