package cars

const defaultProductionRatePerHour int = 221

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	productionRatePerHour := speed * defaultProductionRatePerHour
	return float64(productionRatePerHour) * successRate(speed)
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	defaultProductionRatePerMinute := float64(defaultProductionRatePerHour) / 60.0
	productionRatePerHour := float64(speed) * defaultProductionRatePerMinute
	return int(productionRatePerHour * successRate(speed))
}

// successRate is used to calculate the ratio of an item being created without
// error for a given speed
func successRate(speed int) float64 {
	if speed == 0 {
		return 0.0
	}

	if speed < 5 {
		return 1.0
	}

	if speed >= 9 {
		return 0.77
	}

	return 0.9
}
