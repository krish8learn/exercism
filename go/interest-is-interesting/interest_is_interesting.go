package interest

import (
	"math"
)

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	// panic("Please implement the InterestRate function")
	// panic("Please implement the InterestRate function")
	switch {
	case balance < 0:
		return float32(3.213)
	case balance < 1000:
		return float32(0.5)
	case balance >= 1000 && balance < 5000:
		return float32(1.621)
	case balance >= 5000:
		return float32(2.475)
	}

	return 0
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	// panic("Please implement the Interest function")
	interest := balance * float64(InterestRate(balance)) / 100
	return interest
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	// panic("Please implement the AnnualBalanceUpdate function")
	total := balance + Interest(balance)
	return total
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	// panic("Please implement the YearsBeforeDesiredBalance function")
	// var newBalance float64
	var period int

	for math.Floor(balance) <= math.Floor(targetBalance) {
		balance += Interest(balance)
		period++
	}
	// period++
	// fmt.Println(period, math.RoundToEven(balance), math.RoundToEven(targetBalance))

	return period
}
