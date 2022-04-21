package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	//panic("CalculateWorkingCarsPerHour not implemented")
    successCars := (float64(productionRate) * successRate)/100.0
    return successCars
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	//panic("CalculateWorkingCarsPerMinute not implemented")
    successCarsHour := (float64(productionRate) * successRate)/100.0
    successCarsMinute := int(successCarsHour)/60
    return successCarsMinute
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	//panic("CalculateCost not implemented")
    quotient := carsCount /10
    remainder := carsCount % 10
    cost := quotient*95000 + remainder*10000
    return uint(cost)
}
