// Package weather provides tools to get the weather condition for your city.
package weather

var (
	// CurrentCondition represents a ceratin weather condition
	CurrentCondition string
	// CurrentLocation represents a city location
	CurrentLocation string
)

// Log returns a formated sting with the weather condition of a city
func Log(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
