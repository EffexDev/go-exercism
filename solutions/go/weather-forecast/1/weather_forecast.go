// Package weather This is the weather package. It does weather.
package weather

// CurrentCondition Stores current weather conditions.
var CurrentCondition string
// CurrentLocation Stores the users location.
var CurrentLocation string

// Forecast Looks up the weather in the location and prints it.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
