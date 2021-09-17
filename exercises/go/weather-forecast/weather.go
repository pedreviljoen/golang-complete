// Package weather provides logging functionality for current conditions
package weather

var (
	// CurrentCondition represents the current weather conditions
	CurrentCondition string
	// CurrentCondition represents the current location for the weather conditions
	CurrentLocation string
)

// Log returns a string which describes the current weather conditions at the current location
func Log(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
